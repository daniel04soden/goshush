package cmd

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Prints host name",
	Long:  "Using the gopls host package, we can get the hosts name",
	Run:   whoami,
}

func whoami(cmd *cobra.Command, args []string) {
	hostInfo, _ := host.Info()
	fmt.Println(hostInfo.Hostname)
}

func init() {
	RootCmd.AddCommand(whoamiCmd)
}
