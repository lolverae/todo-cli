package cmd

import (
	"fmt"
	"os"
	"todo-cli/cmd/create"
	"todo-cli/cmd/done"

	"github.com/spf13/cobra"
)

var (
	listName string

	rootCmd = &cobra.Command{
		Use:  "todo-cli [command] [list]",
		Long: "A CLI tool to manage todo lists",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			listName, _ = cmd.Flags().GetString("list")
			if listName != "" {
				CreateTasksFile(listName)
			} else {
				CreateTasksFile("default")
			}
		},
	}
)

func Execute() {
	rootCmd.AddCommand(done.Cmd, create.Cmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVarP(&listName, "list", "l", "", "Name of the task list to modify")
	if e := rootCmd.Execute(); e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", e.Error())
		os.Exit(1)
	}
}
