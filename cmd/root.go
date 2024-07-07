package cmd

import (
	"fmt"
	"github.com/phips4/proj/internal/repo"
	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

// var db *bbolt.DB
var projectRepo repo.ProjectRepo

var rootCmd = &cobra.Command{
	Use:   "proj",
	Short: "proj is a CLI tool to manage projects",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//log.Fatalln("error in root command:", err)
	}
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("error could not get user directory:", err)
	}

	dbPath := filepath.Join(home, "projects.db")

	db, err := bbolt.Open(dbPath, 0666, nil)
	if err != nil {
		log.Fatalln("error could not open database:", err)
	}

	projectRepo = repo.NewBBoltProjectRepository(db)

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(repo.ProjectBucket)
		return err
	})
	if err != nil {
		log.Fatalln("error could not create bucket:", err)
	}

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		all, err := projectRepo.All()
		if err != nil {
			log.Fatalln("error could not get all projects:", err)
			return
		}
		fmt.Println("recently:") // show first four projects
		firstFour := all[:4]
		for _, proj := range firstFour {
			fmt.Println("  "+proj.Name, proj.Execute)
		}

		fmt.Println()
		if len(all) == 0 {
			fmt.Println("you don't have any projects, add a new one like this:")
			fmt.Println("  proj add <name> <description>")
		} else {
			suggestedProj := firstFour[rand.Intn(len(firstFour)-1)]
			fmt.Println("start an existing project:")
			fmt.Println("  proj exec", suggestedProj.Name)
		}
		//fmt.Println()
		//fmt.Println("Available Commands:")
		//for _, c := range cmd.Commands() {
		//	fmt.Printf("  %-20s %s\n", c.UseLine(), c.Short)
		//}
		//
		//fmt.Println()
		//fmt.Println("Flags:")
		//cmd.Flags().VisitAll(func(f *pflag.Flag) {
		//	fmt.Printf("  --%-18s %s (Default: %s)\n", f.Name, f.Usage, f.DefValue)
		//})
		//fmt.Println()
		//fmt.Println("Use \"proj [command] --help\" for more information about a command.")
		//fmt.Println()
		//fmt.Println("Examples:")
		//fmt.Println("  Add a new project:")
		//fmt.Println("    proj add MyProject \"This is my project\" --path=/path/to/project --exec=make")
		//fmt.Println()
		//fmt.Println("  Update project path:")
		//fmt.Println("    proj update MyProject --path=/new/path/to/project")
		//fmt.Println()
		//fmt.Println("  Show details of a project:")
		//fmt.Println("    proj show MyProject")
		//fmt.Println()
		//fmt.Println("  Execute project command:")
		//fmt.Println("    proj exec MyProject")
		//fmt.Println()
	})
}
