package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	updatePath string
	updateExec string
	updateDesc string
)

var updateCmd = &cobra.Command{
	Use:   "update <name or id>",
	Short: "Update project fields",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idOrName := args[0]

		if updatePath != "" {
			if err := projectRepo.UpdatePath(idOrName, updatePath); err != nil {
				return err
			}
			log.Printf("Project %s path updated to %s\n", idOrName, updatePath)
			return nil
		}

		if updateExec != "" {
			// TODO:
			//if err := projectRepo.UpdateExecute(idOrName, updateExec); err != nil {
			//	log.Fatalln("Update Error:", err)
			//	return
			//}
			log.Printf("Project %s execute command updated to %s\n", idOrName, updateExec)
			return nil
		}

		if updateDesc != "" {
			if err := projectRepo.UpdateDescription(idOrName, updateDesc); err != nil {
				return err
			}
			log.Printf("Project %s description updated to %s\n", idOrName, updateDesc)
			return nil
		}

		return nil
	},
}

func init() {
	updateCmd.Flags().StringVar(&updatePath, "path", "", "Update project path")
	updateCmd.Flags().StringVar(&updateExec, "exec", "", "Update project execute command")
	updateCmd.Flags().StringVar(&updateDesc, "desc", "", "Update project description")
	rootCmd.AddCommand(updateCmd)
}
