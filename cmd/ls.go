package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all files in the current directory",
	Long:  "using the io/ioutil package, we list all files found in our current directory",
	Run:   ls,
}

func ls(cmd *cobra.Command, args []string) {
	files, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Print("Directory: ")
		}
		fmt.Println(file.Name())
	}
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
