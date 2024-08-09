package done

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "done [task completed]",
	Short: "Mark a task as completed",
	Long:  `Change the status of an item to completed`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		listContext, _ := cmd.Flags().GetString("list")
		taskTitle := args[0]
		completeTask(taskTitle, listContext)
	},
}

func completeTask(taskTitle string, listContext string) error {
	completeFilePath := ".lists/" + listContext + ".csv"
	file, err := os.Open(completeFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV file: %w", err)
	}

	var newRecords [][]string
	for _, record := range records {
		if len(record) != 2 {
			return fmt.Errorf("invalid CSV format: expected 2 columns")
		}
		if record[0] == taskTitle {
			record[1] = "Complete"
		}
		newRecords = append(newRecords, record)
	}

	tempFile, err := os.CreateTemp("", "modified_*.csv")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	writer := csv.NewWriter(tempFile)
	defer writer.Flush()

	if err := writer.WriteAll(records); err != nil {
		return fmt.Errorf("error writing to temporary file: %w", err)
	}

	if err := os.Rename(tempFile.Name(), completeFilePath); err != nil {
		return fmt.Errorf("error renaming temporary file: %w", err)
	}

	return nil
}
