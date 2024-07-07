package cmd

import (
	"github.com/phips4/proj/internal/repo"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec <name or id>",
	Short: "Execute the command of a project",
	Args:  cobra.ExactArgs(1),
	RunE:  runExecFunc(projectRepo),
}

func runExecFunc(getter repo.ProjectGetter) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		idOrName := args[0]

		proj, err := getter.Get(idOrName)
		if err != nil {
			return err
		}

		cmdExec := exec.Command(proj.Execute)
		if proj.Execute == "" {
			log.Println("proj.Execute is empty, ignoring execution of", proj.Name)
			return nil
		}

		out, err := cmdExec.CombinedOutput()
		if err != nil {
			return err
		}
		log.Println(string(out))
		return nil
	}
}

func init() {
	rootCmd.AddCommand(execCmd)
}
