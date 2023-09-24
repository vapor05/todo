package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/vapor05/todo/src/app"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	RunE:  ListTodo,
}

func ListTodo(cmd *cobra.Command, args []string) error {
	todos, err := app.List(Store)
	if err != nil {
		return err
	}
	fmt.Println("All Todos:")
	fmt.Println("")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Name\tActive\tCreated Date\tDescription")
	for _, t := range todos {
		fmt.Fprintf(w, "%s\t%v\t%s\t%s\t\n", t.Name, t.Active, t.CreatedDate, t.Description)
	}
	w.Flush()
	fmt.Println("")
	return nil
}
