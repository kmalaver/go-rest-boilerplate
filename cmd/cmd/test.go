package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [path]",
	Short: "Run tests.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setting up test environment")
		path := "./pkg/..."
		if len(args) > 0 {
			path = fmt.Sprintf("./pkg/%s/...", args[0])
		}
		fmt.Printf("running tests in %s\n", path)
		command("go test -v -coverprofile=coverage.out -covermode=atomic " + path)
	}}

func init() {
	rootCmd.AddCommand(testCmd)
}
