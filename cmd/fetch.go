package cmd

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Prints user system info",
	Long: `Using the gopsutil library, this command fetches any needed user info
	. It includes a variety of flags to specify this info/format it`,
	Run: fetch,
}

func formatBytes(bytes uint64) string {
	const u = 1024
	if bytes < u {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := int64(u), 0

	for n := bytes / u; n >= u; n /= u {
		div *= u
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func formatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func fetch(cmd *cobra.Command, args []string) {
	Blue := "\033[34m"
	art := Blue + `
                                                                                
        */,****,****,*, ..                             
  .*#&*(     /*,      %*/&*&                           
  %(**.@&    *, &    %,***                            
    ***,**/*,/@@(/********,                            
    ,**,***,*&  /,*,***,**&                            
    ***,*****,*,,*********&                            
    %**,***************,**&                            
.,.%**,******************#.,                          
    #**,*******,*******,**,                            
    ***,******************,                            
    ***,***************,***                            
    %**,*****************,/                            
    /,,*,*,*,*,*,*,*,*,*/                             
    ..#,*************/...                             
    *%.      ..        *.                             

	` + "\033[0m"
	fmt.Println(art)
	info()
}

func init() {
	RootCmd.AddCommand(fetchCmd)
}

func info() {
	vmStat, _ := mem.VirtualMemory()
	diskUsage, _ := disk.Usage("/")
	hostInfo, _ := host.Info()
	cpuInfo, _ := cpu.Info()

	fmt.Println("--- Host Information ---")
	fmt.Printf("OS: %s\n", hostInfo.OS)
	fmt.Printf("Hostname: %s\n", hostInfo.Hostname)
	fmt.Println()

	fmt.Println("--- CPU Information ---")
	fmt.Printf("Model Name: %s\n", cpuInfo[0].ModelName)

	fmt.Println("--- Memory Information ---")
	fmt.Printf("Total: %s\n", formatBytes(vmStat.Total))
	fmt.Printf("Available: %s\n", formatBytes(vmStat.Available))
	fmt.Println()

	fmt.Println("--- Disk Information (/) ---")
	fmt.Printf("Total: %s\n", formatBytes(diskUsage.Total))
	fmt.Printf("Free: %s\n", formatBytes(diskUsage.Free))
	fmt.Println()
}
