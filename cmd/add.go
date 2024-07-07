package cmd

import (
	"github.com/phips4/proj/internal/repo"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	path    string
	execute string
)

var addCmd = &cobra.Command{
	Use:   "add <name> <description>",
	Short: "Add a new project",
	Args:  cobra.ExactArgs(2),
	Run:   runAddFunc(projectRepo),
}

func runAddFunc(adder repo.ProjectAdder) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		name := args[0]
		description := args[1]

		err := adder.Add(name, description, path, execute, []string{})
		if err != nil {
			log.Fatalln("error adding project:", err)
			return
		}

		log.Printf("Project %s added\n", name)
	}
}

func init() {
	addCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {

	})
	addCmd.Flags().StringVar(&path, "path", "", "Path for the project")
	addCmd.Flags().StringVar(&execute, "exec", "", "Execute command for the project")
	rootCmd.AddCommand(addCmd)
}

func itob(v int) []byte {
	return []byte(strconv.Itoa(v))
}
