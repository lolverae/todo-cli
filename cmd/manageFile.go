package cmd

import (
	"log"
	"os"
	"path/filepath"
)

func ManageTasksFile(taskFile string) *os.File {
	// Create the .cache folder if it doesn't exist
	cacheDir := filepath.Dir(taskFile)
	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		log.Printf("error creating cache directory: %s", err)
	}
	_, err = os.Stat(taskFile)
	if os.IsNotExist(err) {
		file, err := os.Create(taskFile)
		log.Printf("Task file %s not found, creating...", taskFile)
		if err != nil {
			log.Printf("Failed to create file")
		}
		return file
	}
	file, err := os.OpenFile(taskFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("error opening file: %s", err)
	}
	return file
}
