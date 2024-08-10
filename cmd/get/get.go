package get

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Gets all tasks",
	Long:  "Gets all tasks with their names and statuses.",
	Args:  cobra.NoArgs,
	RunE:  runCommand,
}

func runCommand(cmd *cobra.Command, args []string) error {
	listContext, _ := cmd.Flags().GetString("list")
	return getTasks(listContext)
}

type Task struct {
	Name   string
	Status string
}

func getTasks(listContext string) error {
	if listContext == "" {
		listContext = "default"
	}
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get home directory: %s", err)
	}
	completeFilePath := filepath.Join(home+"/.lists", listContext+".csv")

	file, err := os.Open(completeFilePath)
	if err != nil {
		return fmt.Errorf("Could not open file %s: %w", completeFilePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Could not read CSV file: %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	for _, task := range tasks {
		if len(task) < 2 {
			return fmt.Errorf("Malformed task record: %v", task)
		}
		fmt.Printf("Name: %s, Status: %s\n", task[0], task[1])
	}

	return nil
}
