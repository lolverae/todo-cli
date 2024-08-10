package create

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "new [task to create]",
	Short: "Create a new task",
	Long:  "Create a new item on the To Do list",
	Args:  cobra.ExactArgs(1),
	RunE:  runCommand,
}

func runCommand(cmd *cobra.Command, args []string) error {
	listContext, _ := cmd.Flags().GetString("list")
	taskTitle := args[0]
	return writeTaskToCSV(taskTitle, listContext)
}

func writeTaskToCSV(taskTitle string, listContext string) error {
	if listContext == "" {
		listContext = "default"
	}
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get home directory: %s", err)
	}
	completeFilePath := filepath.Join(home+"/.lists", listContext+".csv")

	file, err := os.OpenFile(completeFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	task := []string{taskTitle, "Pending"}
	if err := writer.Write(task); err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}
