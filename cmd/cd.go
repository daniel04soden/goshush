package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Changes directory from the current directory to another",
	Long:  `Passes in a single directory or none and from there the shell persists in that new directory`,
	Run:   cd,
}

func cd(cmd *cobra.Command, args []string) {
	newDir := strings.Join(args, " ")
	err := os.Chdir(newDir)
	if err != nil {
		log.Fatal(err)
	}
	cwd,_ := os.Getwd()
	fmt.Println("Changing to directory: ",cwd)
}

func init() {
	RootCmd.AddCommand(cdCmd)
}
