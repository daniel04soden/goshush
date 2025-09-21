package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/prometheus-community/pro-bing"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping takes an ip or a domain and pings that domain",
	Long: `
	Ping is a command based around the posix compliant ping command
	in the Bourne again shell.
	`,
	Run: ping,
}

func ping(cmd *cobra.Command, args []string) {
	pinger,err:=probing.NewPinger(strings.Join(args," "))
	if err!= nil{
		panic(err)
	}

	cancel := make(chan os.Signal, 1)
	signal.Notify(cancel,os.Interrupt)

	go func(){
		for _ = range cancel { // Until cancel
			pinger.Stop()
		}
	}()
	// As it is going
	pinger.OnRecv = func(pkt *probing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	// If packet is a duplicate
	pinger.OnDuplicateRecv = func(pkt *probing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
	}

	// When cancelled
	pinger.OnFinish = func(stats *probing.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
}

func init() {
	RootCmd.AddCommand(pingCmd)
}
