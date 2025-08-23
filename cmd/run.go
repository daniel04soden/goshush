package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runItCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs binaries",
	Long:  "Runs binaries not apart of the current shell",
	Run:   runIt,
}

func runIt(cmd *cobra.Command, args []string) {
	binaryRun := exec.Command(args[0], args[1:]...)
	binaryRun.Stdin = os.Stdin
	binaryRun.Stdout = os.Stdout
	binaryRun.Stderr = os.Stderr

	binaryRun.Run()
}

func init() {
	RootCmd.AddCommand(runItCmd)
}
