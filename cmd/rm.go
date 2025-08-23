package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Rm deletes a file",
	Long:  `Rm takes in a flag of a file to a user and deletes the corresponding file`,
	Run:   rm,
}

func rm(cmd *cobra.Command, args []string) {
	fileName := strings.Join(args, " ")
	err := os.Remove(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted file: ", fileName)
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
