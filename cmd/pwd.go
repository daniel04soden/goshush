package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Prints current directory",
	Long:  "Using the os package, we display the current directory to stdout",
	Run:   pwd,
}

func pwd(cmd *cobra.Command, args []string) {
	currentDir, err := os.Getwd()
	if err != nil {
		print(err)
	} else {
		fmt.Println(currentDir)
	}
}

func init() {
	RootCmd.AddCommand(pwdCmd)
}
