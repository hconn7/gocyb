package cmd

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var subdomainList = []string{
	"www",
	"mail",
	"ftp",
	"admin",
	"api",
	"blog",
	"dev",
	"test",
	"stage",
	"m",
	"app",
	"portal",
	"shop",
	"secure",
	"help",
	"status",
	"support",
	"ns",
	"vpn",
	"intranet",
}
var bruteCmd = &cobra.Command{
	Use:   "brute [URL]",
	Short: "Brute force Subdomains",
	Long:  "Brute force subdomains with common subdomain key words\nTo use your own key words, add a path to a file and the CMD will use the words in it instead",
	Run: func(cmd *cobra.Command, args []string) {
		url := ""

		if len(args) < 1 {
			fmt.Println("Error: URL is required")
			os.Exit(1)
		}
		url = args[0]
		url = urlStrip(url)
		if wordList != "" {
			wordList, err := cmd.Flags().GetString("wordlist")
			if err != nil {
				fmt.Println("Error retrieving wordlist flag:", err)
				os.Exit(1)
				return
			}
			file, err := os.Open(wordList)
			if err != nil {
				fmt.Printf("Failed to open file\nPath:%s", wordList)
				return
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			fmt.Println("----Finding valid SubDomains----")
			for scanner.Scan() {

				subdomains := scanner.Text()

				ips, err := net.LookupHost(subdomains + "." + url)
				if err != nil {
					if dnsErr, ok := err.(*net.DNSError); ok {
						if dnsErr.IsNotFound {
							fmt.Println(err)
						}
					} else {
						fmt.Println("Unkown error:", err)
					}
				}

				for _, ip := range ips {
					fmt.Printf("Trying: %v", ip)
				}
			}

		} else {
			fmt.Println("----Finding valid SubDomains----")
			for _, sub := range subdomainList {

				ips, err := net.LookupHost(sub + "." + url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				for _, ip := range ips {
					fmt.Printf("Trying: %v", ip)

				}

			}

		}

	},
}

func urlStrip(url string) string {
	if strings.HasPrefix(url, "https://") {
		url = strings.TrimPrefix(url, "https://")
	}
	if strings.HasPrefix(url, "http://") {
		url = strings.TrimPrefix(url, "http://")
	}
	if strings.HasPrefix(url, "www") {
		url = strings.TrimPrefix(url, "www")
	}

	return url
}
