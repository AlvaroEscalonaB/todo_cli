package cmd

import (
	"fmt"
	"log"

	"todo_list/internals"

	"github.com/spf13/cobra"
)

var listTasks = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long:  `list all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := internals.GetDatabase()

		if err != nil {
			log.Fatalf("Cannot open or read database %s", err)
		}

		taskRepository := internals.TaskRepository{
			Db: db,
		}
		tasks, err := taskRepository.ListTask()

		if err != nil {
			log.Fatalf("Cannot read the query %s", err)
		}

		fmt.Println(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listTasks)
}
