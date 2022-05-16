package cmd

import (
	"rest/config"

	"github.com/spf13/cobra"
)

// flags
var liveReload = false

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server.",
	Run: func(cmd *cobra.Command, args []string) {
		if liveReload {
			command("go run github.com/cespare/reflex -r .go$ -s -d none go run main.go")
		} else {
			command("go run main.go")
		}
	},
}

func init() {
	serveCmd.Flags().BoolVar(&liveReload, "live-reload", config.Mode.IsDevelopment(), "Enable live reload.")
	rootCmd.AddCommand(serveCmd)
}
