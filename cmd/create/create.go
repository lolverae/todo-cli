package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "new [task to create]",
	Short: "Create a new task",
	Long:  `Create a new item on the To Do list`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createTask()
	},
}

func createTask() error {
	fmt.Print("Creating a new file")
	return nil
}
