package cmd

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [URL]",
	Short: "Scan networks",
	Long:  "Scan networks",
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		for port := 0; port <= 1024; port++ {

			var address string
			if strings.Contains(target, ":") {
				address = fmt.Sprintf("[%s]:%d", target, port)
			} else {
				address = fmt.Sprintf("%s:%d", target, port)
			}
			conn, err := net.DialTimeout("tcp", address, time.Second)
			if err == nil {
				fmt.Printf("Port %d is open\n", port)
				conn.Close()
			}
		}
	},
}
