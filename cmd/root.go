package cmd

import (
	"fmt"
	"os"
	"todo-cli/cmd/create"
	"todo-cli/cmd/done"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "bin [command]",
	Long:          "",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
	},
}

func Execute() {
	rootCmd.AddCommand(done.Cmd, create.Cmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if e := rootCmd.Execute(); e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", e.Error())
		os.Exit(1)
	}
}
