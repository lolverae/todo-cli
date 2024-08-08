package cmd

import (
	"log"
	"os"
	"path/filepath"
)

func CreateTasksContext(taskFile string) *os.File {
	completeFilePath := ".lists/" + taskFile + ".csv"
	cacheDir := filepath.Dir(completeFilePath)
	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		log.Printf("error creating lists directory: %s", err)
	}

	_, err = os.Stat(completeFilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(completeFilePath)
		log.Printf("Task file %s not found, creating...", taskFile)
		if err != nil {
			log.Printf("Failed to create file")
		}
		return file
	}

	file, err := os.OpenFile(completeFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("error opening file: %s", err)
	}
	return file
}
