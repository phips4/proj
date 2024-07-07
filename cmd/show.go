package cmd

import (
	"github.com/phips4/proj/internal/repo"
	"log"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <name or id>",
	Short: "Show the details of a project",
	Args:  cobra.ExactArgs(1),
	RunE:  runShowFunc(projectRepo),
}

func runShowFunc(getter repo.ProjectGetter) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		idOrName := args[0]

		proj, err := getter.Get(idOrName)
		if err != nil {
			return err
		}

		log.Printf("project: %s\n,  desc: %s\n,  path: %s\n,  exec %s\n", proj.Name, proj.Description, proj.Path, proj.Execute)
		return nil
	}
}

func init() {
	rootCmd.AddCommand(showCmd)
}
