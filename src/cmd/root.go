package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vapor05/todo/src/storage"
)

var (
	Store    *storage.JSONStorage
	DataFile string
)

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
	RootCmd.PersistentPreRunE = SetupStorage
}

func SetupStorage(cmd *cobra.Command, args []string) error {
	var err error
	Store, err = storage.NewJSONStorage(DataFile)
	if err != nil {
		return err
	}
	return nil
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
