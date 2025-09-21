package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var all bool  
var rmdirCmd = &cobra.Command{
	Use:   "rmdir",
	Short: "rmdir removes a directory",
	Long:  `rmdir takes in a flag for the name from a user to delete a directory corresponding to that name`,
	Run:   rmdir,
}

func rmdir(cmd *cobra.Command, args []string) {
	dirName := strings.Join(args, " ")
	if !all {
		if err := os.Remove(dirName); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := os.RemoveAll(dirName); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Removed directory", dirName)
}

func init() {
	rmdirCmd.Flags().BoolVarP(&all, "all", "f", false, "Ensures all files even if non empty are deleted")
	RootCmd.AddCommand(rmdirCmd)
}
