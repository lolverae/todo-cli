package get

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Gets all tasks",
	Long:  `Gets all tasks name and its status`,
	Run: func(cmd *cobra.Command, args []string) {
		listContext, _ := cmd.Flags().GetString("list")
		getTasks(listContext)
	},
}

type Task struct {
	Name   string
	Status string
}

func getTasks(listContext string) error {
	completeFilePath := ".lists/" + listContext + ".csv"
	file, err := os.Open(completeFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var result []Task
	for _, record := range records[:] {
		result = append(result, Task{
			Name:   record[0],
			Status: record[1],
		})
	}

	for _, record := range result {
		fmt.Printf("Name: %s, Status: %s\n", record.Name, record.Status)
	}
	return nil
}
