package get

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/table"
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
	tasksList, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Could not read CSV file: %w", err)
	}

	// Map CSV to Task structs
	var tasks []Task
	for _, record := range tasksList {
		if len(record) < 2 {
			return fmt.Errorf("malformed task record: %v", record)
		}
		tasks = append(tasks, Task{
			Name:   record[0],
			Status: record[1],
		})
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	// Setup output as table using pretty print
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Task", "Status"})
	for _, task := range tasks {
		t.AppendRows([]table.Row{
			{task.Name, task.Status},
		})
	}

	t.SetStyle(table.StyleBold)
	t.Render()

	return nil
}
