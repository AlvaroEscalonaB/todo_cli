package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of todo_list application",
	Long:  `Shows the version of todo_list application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version of app is 0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
