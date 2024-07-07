package cmd

import (
	"github.com/phips4/proj/internal/repo"
	"log"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <name or id>",
	Short: "Delete a project",
	Args:  cobra.ExactArgs(1),
	Run:   runDeleteFunc(projectRepo),
}

func runDeleteFunc(deleter repo.ProjectDeleter) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		idOrName := args[0]

		if err := deleter.Delete(idOrName); err != nil {
			log.Fatalln("error deleting project", err)
			return
		}

		log.Printf("Project %s deleted\n", idOrName)
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
