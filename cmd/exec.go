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
	Run:   runExecFunc(projectRepo),
}

func runExecFunc(getter repo.ProjectGetter) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		idOrName := args[0]

		proj, err := getter.Get(idOrName)
		if err != nil {
			log.Fatalln("Couldn't find project:", err)
		}

		cmdExec := exec.Command(proj.Execute)
		if proj.Execute != "" {
			out, err := cmdExec.CombinedOutput()
			if err != nil {
				log.Fatalln("Error executing:", err)
				return
			}
			log.Println(string(out))
		} else {
			log.Println("do command in project:", proj.Name)
		}
	}
}

func init() {
	rootCmd.AddCommand(execCmd)
}
