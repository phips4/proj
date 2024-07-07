package cmd

import (
	"github.com/phips4/proj/internal/repo"
	"log"

	"github.com/spf13/cobra"
)

var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Manage project labels",
}

var addLabelCmd = &cobra.Command{
	Use:   "add <name or id> <label>",
	Short: "Add a label to a project",
	Args:  cobra.ExactArgs(2),
	Run:   runAddLabelFunc(projectRepo),
}

func runAddLabelFunc(plu repo.ProjectLabelManager) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		idOrName := args[0]
		label := args[1]

		err := plu.AddLabel(idOrName, label)
		if err != nil {
			log.Fatalln("Error adding label:", err)
			return
		}

		log.Printf("Label added to project %s\n", idOrName)
	}
}

var removeLabelCmd = &cobra.Command{
	Use:   "remove <name or id> <label>",
	Short: "Remove a label from a project",
	Args:  cobra.ExactArgs(2),
	Run:   runRemoveLabelFunc(projectRepo),
}

func runRemoveLabelFunc(plu repo.ProjectLabelManager) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		idOrName := args[0]
		label := args[1]

		if err := plu.RemoveLabel(idOrName, label); err != nil {
			log.Fatalln("Error removing label:", err)
			return
		}

		log.Printf("Label removed from project %s\n", idOrName)
	}
}

func init() {
	labelCmd.AddCommand(addLabelCmd)
	labelCmd.AddCommand(removeLabelCmd)
	rootCmd.AddCommand(labelCmd)
}
