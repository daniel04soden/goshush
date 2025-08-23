package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Prints Contents of a file",
	Long:  `Cat is a command that loops through a file passed `,
	Run:   cat,
}

func readFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("file not found")
	}
	return string(content)
}

func cat(cmd *cobra.Command, args []string) {
	content := readFile(strings.Join(args, " "))
	fmt.Println(content)
}

func init() {
	RootCmd.AddCommand(catCmd)
}
