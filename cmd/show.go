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
	Run:   runShowFunc(projectRepo),
}

func runShowFunc(getter repo.ProjectGetter) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		idOrName := args[0]

		proj, err := getter.Get(idOrName)
		if err != nil {
			log.Fatalln("Couldn't find project:", err)
			return
		}
		log.Printf("project: %s\n,  desc: %s\n,  path: %s\n,  exec %s\n", proj.Name, proj.Description, proj.Path, proj.Execute)
	}
}

func init() {
	rootCmd.AddCommand(showCmd)
}
