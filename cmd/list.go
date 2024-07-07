package cmd

import (
	"fmt"
	"github.com/phips4/proj/internal/repo"
	"github.com/spf13/cobra"
	"log"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all projects",
	Args:  cobra.ExactArgs(0),
	RunE:  runListFunc(projectRepo),
}

func runListFunc(getter repo.ProjectGetter) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		projects, err := getter.All()
		if err != nil {
			return err
		}

		for i, proj := range projects {
			fmt.Println(i, "project:", proj.Name, proj.Description)
		}
		log.Println("shown", len(projects), "projects")
		return nil
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
