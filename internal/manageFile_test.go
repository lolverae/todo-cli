package internal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateTasksFile(t *testing.T) {
	testCases := []struct {
		name        string
		taskFile    string
		expectError bool
		expectFile  bool
		expectDir   bool
	}{
		{
			name:        "Create New File",
			taskFile:    "testfile",
			expectError: false,
			expectFile:  true,
			expectDir:   true,
		},
		{
			name:        "File Already Exists",
			taskFile:    "existingfile",
			expectError: false,
			expectFile:  true,
			expectDir:   true,
		},
		{
			name:        "Invalid File Name",
			taskFile:    string([]rune{0x007F}), // Invalid file name character
			expectError: true,
			expectFile:  false,
			expectDir:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare the environment
			testFilePath := filepath.Join(".lists", tc.taskFile+".csv")

			// Clean up before test
			os.RemoveAll(".lists")

			// Call the function
			err := CreateTasksFile(tc.taskFile)

			// Check for errors
			if (err != nil) != tc.expectError {
				t.Errorf("CreateTasksFile() error = %v, expectError %v", err, tc.expectError)
				return
			}

			// Check if file and directory exist
			fileInfo, err := os.Stat(testFilePath)
			if (err != nil && !os.IsNotExist(err)) || (err == nil && !tc.expectFile) {
				t.Errorf("Expected file existence %v, got %v", tc.expectFile, err == nil)
			}

			if !tc.expectFile && fileInfo != nil {
				t.Errorf("File %s should not exist", testFilePath)
			}

			dirInfo, err := os.Stat(filepath.Dir(testFilePath))
			if (err != nil && !os.IsNotExist(err)) || (err == nil && !tc.expectDir) {
				t.Errorf("Expected directory existence %v, got %v", tc.expectDir, err == nil)
			}

			if !tc.expectDir && dirInfo != nil {
				t.Errorf("Directory %s should not exist", filepath.Dir(testFilePath))
			}

			// Clean up after test
			os.RemoveAll(".lists")
		})
	}
}
