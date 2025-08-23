package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var mkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "mkdir creates a directory",
	Long:  `mkdir takes in a flag for the name from a user to create a directory of that name`,
	Run:   mkdir,
}

func mkdir(cmd *cobra.Command, args []string) {
	dirName := strings.Join(args, " ")
	if err := os.Mkdir(dirName, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created directory", dirName)
}

func init() {
	RootCmd.AddCommand(mkdirCmd)
}
