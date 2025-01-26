package cmd

import (
	"fmt"
	"log"
	"todo_list/internals"

	"github.com/spf13/cobra"
)

var addTodo = &cobra.Command{
	Use:   "add [your todo] --description",
	Short: "add",
	Long:  `add`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			log.Fatalf("Error on get flag 'description' %v", description)
			return
		}
		db, err := internals.GetDatabase()

		if err != nil {
			log.Fatalf("Cannot open or read database %s", err)
		}

		taskRepository := internals.TaskRepository{
			Db: db,
		}
		task := internals.InsertTask{
			Name:        args[0],
			Description: description,
		}
		tasks, err := taskRepository.NewTask(task)

		if err != nil {
			log.Fatalf("Cannot read the query %s", err)
		}

		fmt.Println(tasks)
	},
}

func init() {
	rootCmd.AddCommand(addTodo)
	addTodo.PersistentFlags().String("description", "", "Additional description for the task")
}
