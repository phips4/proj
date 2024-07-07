package cmd

import (
	"github.com/phips4/proj/internal/repo"
	"github.com/spf13/cobra"
	"log"
)

var (
	path    string
	execute string
)

var addCmd = &cobra.Command{
	Use:   "add <name> <description>",
	Short: "Add a new project",
	Args:  cobra.ExactArgs(2),
	RunE:  runAddFunc(projectRepo),
}

func runAddFunc(adder repo.ProjectAdder) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		name := args[0]
		description := args[1]

		err := adder.Add(name, description, path, execute, []string{})
		if err != nil {
			return err
		}

		log.Printf("Project %s added\n", name)
		return nil
	}
}

func init() {
	addCmd.Flags().StringVar(&path, "path", "", "Path for the project")
	addCmd.Flags().StringVar(&execute, "exec", "", "Execute command for the project")
	rootCmd.AddCommand(addCmd)
}
