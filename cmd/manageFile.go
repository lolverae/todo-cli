package cmd

import (
	"log"
	"os"
	"path/filepath"
)

var tasksFile *os.File

func CreateTasksFile(taskFile string) error {
	completeFilePath := ".lists/" + taskFile + ".csv"
	cacheDir := filepath.Dir(completeFilePath)
	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		log.Printf("error creating lists directory: %s", err)
		return err
	}

	_, err = os.Stat(completeFilePath)
	if os.IsNotExist(err) {
		_, err := os.Create(completeFilePath)
		if err != nil {
			log.Printf("Failed to create file")
			return err
		}
		return nil
	}
	return nil
}
