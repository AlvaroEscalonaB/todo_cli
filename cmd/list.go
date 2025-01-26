package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"todo_list/internals"

	"github.com/spf13/cobra"
)

const (
	id          = "id"
	name        = "name"
	description = "description"
	date        = "date"
	completed   = "completed"
)

var listTasks = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long:  `list all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		completedFlag, err := cmd.Flags().GetBool("completed")
		if err != nil {
			log.Fatalf("Issue obtaining the flag of completed: %s", err)
		}

		db, err := internals.GetDatabase()

		if err != nil {
			log.Fatalf("Cannot open or read database %s", err)
		}

		taskRepository := internals.TaskRepository{
			Db: db,
		}
		tasks, err := taskRepository.ListTask(completedFlag)

		if err != nil {
			log.Fatalf("Cannot read the query %s", err)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 2, 3, ' ', 0)
		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\n", id, name, description, date, completed)
		for _, task := range tasks {
			fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%t\n", task.Id, task.Name, task.Description, task.Date, task.Completed)
		}

		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listTasks)
	listTasks.PersistentFlags().Bool("completed", false, "Filter the tasks if are completed")
}
