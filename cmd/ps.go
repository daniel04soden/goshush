package cmd

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/process"
	"github.com/spf13/cobra"
)
var (
	nall bool
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "Information about current running process",
	Long:  `Using the gopsutil package we are able to extract information relating to running process(es) and display this information to the user`,
	Run:   ps,
}

func ps(cmd *cobra.Command, args []string) {
	processes,_ := process.Processes()
	if nall{
		firstProcess := processes[0]
		name, _ := firstProcess.Name()
		memory,_ := firstProcess.MemoryPercent()
		fmt.Printf("ID: %d | Process: %s | Memory Percentage used: %f \n", firstProcess.Pid, name,memory)
		return
	}
		for _, proc := range processes {
			name, _ := proc.Name()
			memory,_ := proc.MemoryPercent()
			fmt.Printf("ID: %d | Process: %s | Memory Percentage used: %f \n", proc.Pid, name,memory)
		}
}

func init() {
	psCmd.Flags().BoolVarP(&nall, "nall", "n", false, "Does not display all running processes")
	RootCmd.AddCommand(psCmd)
}
