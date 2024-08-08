package done

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "done [task completed]",
	Short: "Mark a task as completed",
	Long:  `Change the status of an item to completed`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		completeTask()
	},
}

func completeTask() error {
	fmt.Println("Task marked as completed!")
	return nil
}
