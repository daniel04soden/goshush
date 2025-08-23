package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
    "github.com/shirou/gopsutil/v4/mem"
    "github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
)

var fetchCmd = &cobra.Command{
	Use:"fetch",
	Short:"Prints user system info",
	Long : `Using the gopsutil library, this command fetches any needed user info
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

func fetch(cmd *cobra.Command, args []string){
	var art string = `
                                                                                
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

	`
	fmt.Println(art)
	info()
}

func init(){
	rootCmd.AddCommand(fetchCmd)
}



func info() {
    vmStat, _ := mem.VirtualMemory()
	diskUsage,_:= disk.Usage("/")
	hostInfo,_ := host.Info()
	cpuInfo,_ := cpu.Info()


	fmt.Println("--- Host Information ---")
	fmt.Printf("OS: %s\n", hostInfo.OS)
	fmt.Printf("Hostname: %s\n", hostInfo.Hostname)
	fmt.Println()

	fmt.Println("--- CPU Information ---")
	fmt.Printf("Model Name: %s\n", cpuInfo[0].ModelName) // Note: This assumes cpuInfo is not empty

	fmt.Println("--- Memory Information ---")
	fmt.Printf("Total: %s\n", formatBytes(vmStat.Total))
	fmt.Printf("Available: %s\n", formatBytes(vmStat.Available))
	fmt.Println()

	fmt.Println("--- Disk Information (/) ---")
	fmt.Printf("Total: %s\n", formatBytes(diskUsage.Total))
	fmt.Printf("Free: %s\n", formatBytes(diskUsage.Free))
	fmt.Println()
}

func verbose(){
    vmStat, _ := mem.VirtualMemory()
	diskUsage,_:= disk.Usage("/")
	hostInfo,_ := host.Info()
	cpuInfo,_ := cpu.Info()


	fmt.Println("--- Host Information ---")
	fmt.Printf("Hostname: %s\n", hostInfo.Hostname)
	fmt.Printf("OS: %s\n", hostInfo.OS)
	fmt.Printf("Kernel Version: %s\n", hostInfo.KernelVersion)
	fmt.Printf("Uptime: %s\n", formatDuration(time.Duration(hostInfo.Uptime)*time.Second))
	fmt.Printf("Boot Time: %s\n", time.Unix(int64(hostInfo.BootTime), 0).Format(time.RFC1123))
	fmt.Println()

	fmt.Println("--- CPU Information ---")
	fmt.Printf("Logical Cores: %d\n", len(cpuInfo))
	fmt.Printf("Model Name: %s\n", cpuInfo[0].ModelName) // Note: This assumes cpuInfo is not empty

	fmt.Println("--- Memory Information ---")
	fmt.Printf("Total: %s\n", formatBytes(vmStat.Total))
	fmt.Printf("Available: %s\n", formatBytes(vmStat.Available))
	fmt.Printf("Used: %s\n", formatBytes(vmStat.Used))
	fmt.Printf("Used Percent: %.2f%%\n", vmStat.UsedPercent)
	fmt.Println()

	fmt.Println("--- Disk Information (/) ---")
	fmt.Printf("Total: %s\n", formatBytes(diskUsage.Total))
	fmt.Printf("Free: %s\n", formatBytes(diskUsage.Free))
	fmt.Printf("Used: %s\n", formatBytes(diskUsage.Used))
	fmt.Printf("Used Percent: %.2f%%\n", diskUsage.UsedPercent)
	fmt.Println()
}
