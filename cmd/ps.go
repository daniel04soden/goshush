package cmd

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/process"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "Information about running processes",
	Long:  `Using the gopsutil package we are able to extract information relating to all running processes and display this information to the user`,
	Run:   ps,
}

func ps(cmd *cobra.Command, args []string) {
	procs, _ := process.Processes()
    for _, proc := range procs {
        name, _ := proc.Name()
        fmt.Printf("ID: %d | Process: %s\n", proc.Pid, name)
    }
}

func init() {
	RootCmd.AddCommand(psCmd)
}
