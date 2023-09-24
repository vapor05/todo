package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var DataFile string

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "manage a todo list",
	Long: `Keep track of tasks with this todo cli.

Add new tasks, edit tasks, and complete tasks using the provided commands`,
}

func init() {
	RootCmd.PersistentFlags().StringVar(
		&DataFile, "data-file", "todos.json",
		"name of json file used to store todo data",
	)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
