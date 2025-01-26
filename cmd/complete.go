package cmd

import (
	"fmt"
	"log"
	"todo_list/internals"

	"github.com/spf13/cobra"
)

var completeTodo = &cobra.Command{
	Use:   "complete --id",
	Short: "add",
	Long:  `add`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil || id == -1 {
			log.Fatalf("Error on get flag 'id' %v", id)
			return
		}

		db, err := internals.GetDatabase()
		if err != nil {
			log.Fatalf("Cannot open or read database %s", err)
		}

		taskRepository := internals.TaskRepository{
			Db: db,
		}

		tasks, err := taskRepository.CompleteTask(id)

		if err != nil {
			log.Fatalf("Cannot read the query %s", err)
		}

		fmt.Println(tasks)
	},
}

func init() {
	rootCmd.AddCommand(completeTodo)
	completeTodo.PersistentFlags().Int("id", -1, "Id that that want to complete")
}
