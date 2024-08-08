package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func Execute() {
	cmdNew := &cobra.Command{
		Use:   "new [task to create]",
		Short: "Create a new task",
		Long:  `Create a new item on the To Do list`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("New: " + strings.Join(args, " "))
		},
	}

	cmdDone := &cobra.Command{
		Use:   "done [task completed]",
		Short: "Mark a task as completed",
		Long:  `Change the status of an item to completed`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	rootCmd := &cobra.Command{Use: "todo-cli"}
	rootCmd.AddCommand(cmdNew, cmdDone)
	rootCmd.Execute()
}
