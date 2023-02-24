package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Shemistan/EnglishBot/cmd/migrate/initdb"
	"github.com/Shemistan/EnglishBot/cmd/migrate/up"
)

var rootCmd = &cobra.Command{
	Short: "admin",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	initDBCmd := initdb.Command()
	rootCmd.AddCommand(initDBCmd)

	upCmd, err := up.Command()
	if err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(upCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
