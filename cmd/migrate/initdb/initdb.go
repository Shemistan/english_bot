package initdb

import (
	"database/sql"
	"fmt"
	"strings"

	// required to user with cockroachdb
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// prepared statement params is not working here
	// Could not execute query CREATE DATABASE IF NOT EXISTS $1; with params vw_page_test: pq: syntax error at or near "1"
	createDBQuery        = "CREATE DATABASE IF NOT EXISTS %s;"
	createUserQuery      = "CREATE USER IF NOT EXISTS %s;"
	grantPrivilegesQuery = "GRANT ALL ON DATABASE %s TO %s;"
)

// Command creates initdb command
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "initdb",
		Short: "Command creates database and user",
		Long:  "Command creates database and user, grant privileges to user on database",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return errors.Wrap(err, "failed to bind flags")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			dsn := viper.GetString("sql_connection_string")
			dbName := viper.GetString("dbname")
			dbUser := viper.GetString("dbuser")

			// FIXME
			if strings.HasPrefix(dsn, "cockroachdb://") {
				dsn = strings.Replace(dsn, "cockroachdb://", "postgres://", 1)
			}

			l := log.WithField("dbName", dbName).
				WithField("dbUser", dbUser)

			l.Info("Creating db and user")

			db, err := sql.Open("postgres", dsn)
			if err != nil {
				return errors.Wrapf(err, "Could not open sql connection to %s", dsn)
			}
			defer db.Close()
			tx, err := db.Begin()
			if err != nil {
				return errors.Wrapf(err, "Could not start transaction")
			}
			rollBack := func() {
				err := tx.Rollback()
				if err != nil {
					l.Error(errors.Wrap(err, "Could not rollback transaction"))
				}
			}

			if _, err := tx.Exec(fmt.Sprintf(createDBQuery, dbName)); err != nil {
				rollBack()
				return errors.Wrapf(err, "Could not execute query %s with params %s", createDBQuery, dbName)
			}

			if _, err := tx.Exec(fmt.Sprintf(createUserQuery, dbUser)); err != nil {
				rollBack()
				return errors.Wrapf(err, "Could not execute query %s with params %s", createUserQuery, dbUser)
			}

			if _, err := tx.Exec(fmt.Sprintf(grantPrivilegesQuery, dbName, dbUser)); err != nil {
				rollBack()
				return errors.Wrapf(err, "Could not execute query %s with params %s %s", grantPrivilegesQuery, dbName, dbUser)
			}

			if err := tx.Commit(); err != nil {
				return errors.Wrapf(err, "Could not commit transaction")
			}

			l.Info("DB and user creation finished")

			return nil
		},
	}

	cmd.Flags().StringP("sql_connection_string", "c", "",
		"Database data source name (e.g psql://user:pass@host:1232)")
	cmd.Flags().StringP("dbname", "d", "",
		"Database name to create (e.g vw_page")
	cmd.Flags().StringP("dbuser", "u", "",
		"User to create to user with database (e.g. vw_page_roach)")
	if err := cmd.MarkFlagRequired("sql_connection_string"); err != nil {
		log.Fatalf("failed to mark option \"sql_connection_string\" as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("dbname"); err != nil {
		log.Fatalf("failed to mark option \"dbname\" as required: %v", err)
	}
	if err := cmd.MarkFlagRequired("dbuser"); err != nil {
		log.Fatalf("failed to mark option \"user\" as required: %v", err)
	}

	return cmd
}
