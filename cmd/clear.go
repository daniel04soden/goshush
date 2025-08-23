package cmd

import (
	"fmt"
	"runtime"

	"github.com/inancgumus/screen"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Lists all files in the current directory",
	Long:  "using the io/ioutil package, we list all files found in our current directory",
	Run:   clear,
}

func clear(cmd *cobra.Command, args []string) {
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		fmt.Println("\033[2J")
	} else {
		screen.Clear()
	}
}

func init() {
	RootCmd.AddCommand(clearCmd)
}
