package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echo takes in from standard in and returns to stdout",
	Long: `
	Echo is a command based around the posix compliant echo command
	in the Bourne again shell.
	`,
	Run: echo,
}

func echo(cmd *cobra.Command, args []string) {
	fmt.Println(strings.Join(args, " "))
}

func init() {
	RootCmd.AddCommand(echoCmd)
}
