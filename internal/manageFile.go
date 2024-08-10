package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// CreateTasksFile creates a new CSV file for tasks in a specified directory.
// It ensures the directory structure exists before creating the file.
func CreateTasksFile(taskFile string) error {
	if !isValidTaskFile(taskFile) {
		return fmt.Errorf("invalid taskFile: must contain only letters and underscores")
	}
	// Construct the complete file path
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Failed to get home directory: %s", err)
	}
	completeFilePath := filepath.Join(home+"/.lists", taskFile+".csv")

	// Ensure the directory exists
	listsDir := filepath.Dir(completeFilePath)
	err = os.MkdirAll(listsDir, 0777)
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

func isValidTaskFile(taskFile string) bool {
	// Regular expression to match only letters and underscores
	validFileRegex := regexp.MustCompile(`^[a-zA-Z_]+$`)
	return validFileRegex.MatchString(taskFile)
}
