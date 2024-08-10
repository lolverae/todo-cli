package internal

import (
	"log"
	"os"
	"path/filepath"
)

// CreateTasksFile creates a new CSV file for tasks in a specified directory.
// It ensures the directory structure exists before creating the file.
func CreateTasksFile(taskFile string) error {
	// Construct the complete file path
	completeFilePath := filepath.Join(".lists", taskFile+".csv")

	// Ensure the directory exists
	cacheDir := filepath.Dir(completeFilePath)
	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		log.Printf("error creating lists directory: %s", err)
		return err
	}

	// Check if the file already exists
	if _, err := os.Stat(completeFilePath); os.IsNotExist(err) {
		// Create the file if it doesn't exist
		if _, err := os.Create(completeFilePath); err != nil {
			log.Printf("Failed to create file: %s", err)
			return err
		}
	}
	return nil
}
