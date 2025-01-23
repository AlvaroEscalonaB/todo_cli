package cmd

import (
	"fmt"
	"log"

	"todo_list/internals"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo_list",
	Short: "todo_list is a CLI tool to register activities",
	Long:  "todo_list is a CLI tool to register activities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to todo_list CLI")
	},
}

func Execute() {
	if err := internals.CreateDatabase(); err != nil {
		log.Fatalf("Error on database creation %v", err)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
