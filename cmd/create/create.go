package create

import (
	"encoding/csv"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "new [task to create]",
	Short: "Create a new task",
	Long:  `Create a new item on the To Do list`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		listContext, _ := cmd.Flags().GetString("list")
		taskTitle := args[0]
		writeTaskToCSV(taskTitle, listContext)
	},
}

type Task struct {
	Name   string
	Status string
}

func writeTaskToCSV(taskTitle string, listContext string) error {
	completeFilePath := ".lists/" + listContext + ".csv"
	file, err := os.OpenFile(completeFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var task Task
	task.Name = taskTitle
	task.Status = "Pending"
	err = writer.Write([]string{task.Name, task.Status})
	if err != nil {
		return err
	}

	return nil
}
