package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var wordList string
var rootCmd = &cobra.Command{
	Use:   "gocyb",
	Short: "CyberSecurity CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gocyb! Use the 'brute' command for subdomain brute force.")
	},
}

func init() {
	rootCmd.AddCommand(bruteCmd)
	bruteCmd.Flags().StringVarP(&wordList, "wordlist", "w", "", "Path to custom wordlist file")
	rootCmd.AddCommand(scanCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
