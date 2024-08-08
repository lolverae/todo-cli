package main

import (
	// "todo-cli/cmd"
	"todo-cli/internal"
)

func main() {
	// cmd.Execute()
	taskFile := "./.cache/tasks.csv"
	internal.ManageTasksFile(taskFile)
}
