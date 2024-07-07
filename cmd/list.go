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
	Run:   runListFunc(projectRepo),
}

func runListFunc(getter repo.ProjectGetter) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		projects, err := getter.All()
		if err != nil {
			log.Fatal("error getting all projects:", err)
		}

		for i, proj := range projects {
			fmt.Println(i, "project:", proj.Name, proj.Description)
		}
		log.Println("shown", len(projects), "projects")
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
