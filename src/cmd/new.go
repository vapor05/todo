package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new todo",
	RunE:  NewTodo,
}

func NewTodo(cmd *cobra.Command, args []string) error {
	todo := strings.Join(args, " ")
	fmt.Printf("creating a new todo, %s!\n", todo)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to create new todo: %w", err)
	}
	description := strings.TrimSuffix(input, "\n")
	fmt.Println("provide a description for your new todo:")
	fmt.Printf("created new todo, %s, with description: %s\n", todo, description)
	return nil
}
