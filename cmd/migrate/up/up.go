package up

import (
	"github.com/golang-migrate/migrate/v4"
	// CockroachDB installation
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
)

func Command() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "up",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			errBind := viper.BindPFlags(cmd.Flags())
			if errBind != nil {
				return errors.Wrap(errBind, "failed to bind flags")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				migrationsPath = viper.GetString("path")
				dbPath         = viper.GetString("db")
			)

			m, err := migrate.New(migrationsPath, dbPath)
			if err != nil {
				return err
			}

			err = m.Up()
			if err != nil {
				if errors.Is(err, migrate.ErrNoChange) {
					return nil
				}

				return err
			}

			return nil
		},
	}

	cmd.Flags().String("path", "", "migrations path (file://migrations")
	cmd.Flags().String("db", "", "DB path (cockroachdb://root:@localhost:26257/cherrity?sslmode=disable)")

	err := multierr.Combine(
		cmd.MarkFlagRequired("path"),
		cmd.MarkFlagRequired("db"),
	)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
