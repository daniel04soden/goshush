package cmd

import (
	"fmt"
	"runtime"

	"github.com/inancgumus/screen"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the terminal screen",
	Long:  "Depending on the OS being used, it either clears the screen using a fmt ascii code or uses the screen package by spf13",
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
