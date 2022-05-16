package cmd

import (
	"rest/tools"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [command]",
	Short: "Manage migrations.",
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Run all pending migrations.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.NewMigrator("./db/migrations").Up()
	}}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Roll back the last migration.",
	Run: func(cmd *cobra.Command, args []string) {
		tools.NewMigrator("./db/migrations").Down()
	}}

var migrateNewCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new migration file.",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tools.NewMigrator("./db/migrations").New(args[0])
	}}

func init() {
	migrateCmd.AddCommand(migrateUpCmd, migrateDownCmd, migrateNewCmd)
	rootCmd.AddCommand(migrateCmd)
}
