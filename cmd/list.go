package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"todo_list/internals"

	"github.com/fatih/color"
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

		blue := color.New(color.FgCyan).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()
		writer := tabwriter.NewWriter(os.Stdout, 0, 2, 3, ' ', 0)

		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\n", blue(id), blue(name), blue(description), blue(date), blue(completed))

		for _, task := range tasks {
			if task.Completed {
				fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\n", green(task.Id), green(task.Name), green(task.Description), green(task.Date.Format("2006-01-02 15:04:05")), green(task.Completed))
			} else {
				fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\n", red(task.Id), red(task.Name), red(task.Description), red(task.Date.Format("2006-01-02 15:04:05")), red(task.Completed))
			}
		}

		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listTasks)
	listTasks.PersistentFlags().Bool("completed", false, "Filter the tasks if are completed")
}
