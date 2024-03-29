package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vapor05/todo/src/app"
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
	name := strings.Join(args, " ")
	fmt.Printf("creating a new todo, %s!\n", name)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("provide a description for your new todo:")
	input, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to create new todo: %w", err)
	}
	description := strings.TrimSuffix(input, "\n")
	err = app.New(name, description, Store)
	if err != nil {
		return err
	}
	fmt.Printf("created new todo, %s, with description: %s\n", name, description)
	return nil
}
