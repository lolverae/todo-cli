package done

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "done [task completed]",
	Short: "Mark a task as completed",
	Long:  `Change the status of an item to completed`,
	Args:  cobra.ExactArgs(1),
	RunE:  runCommand,
}

func runCommand(cmd *cobra.Command, args []string) error {
	listContext, _ := cmd.Flags().GetString("list")
	taskTitle := args[0]
	return completeTask(taskTitle, listContext)
}

func completeTask(taskTitle string, listContext string) error {
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
		return fmt.Errorf("could not open file %s: %w", completeFilePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV file: %w", err)
	}

	var newRecords [][]string
	taskFound := false
	for _, record := range records {
		if len(record) != 2 {
			return fmt.Errorf("invalid CSV format: expected 2 columns, found %d", len(record))
		}
		if record[0] == taskTitle {
			record[1] = "Complete"
			taskFound = true
		}
		newRecords = append(newRecords, record)
	}

	if !taskFound {
		return fmt.Errorf("task %q not found", taskTitle)
	}

	// Create temp file to store the new file with new statuses
	tempFile, err := os.CreateTemp("", "modified_*.csv")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Write on new temp file
	writer := csv.NewWriter(tempFile)
	if err := writer.WriteAll(newRecords); err != nil {
		return fmt.Errorf("error writing to temporary file: %w", err)
	}
	writer.Flush()

	// Replace old list with updated list
	if err := os.Rename(tempFile.Name(), completeFilePath); err != nil {
		return fmt.Errorf("error renaming temporary file: %w", err)
	}

	return nil
}
