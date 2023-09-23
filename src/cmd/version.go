package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get version of the todo cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
