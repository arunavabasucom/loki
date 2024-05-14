package cmd

import (
	"fmt"
	"net"
	"os"
	"github.com/spf13/cobra"

)

var rootCmd = &cobra.Command{
	Use:   "loki.projects.arunavabasu.com",
	Short: "Modern CLI for DNS lookups",
	Long: `loki.projects.arunavabasu.com is a modern CLI for DNS lookups.
	This application allows you to perform DNS lookups for a given hostname.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please provide a domain name")
			return
		}

		domain := args[0]
		ips, err := lookupIP(domain)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("IP addresses for %s:\n", domain)
		for _, ip := range ips {
			fmt.Println(ip)
		}
	},
}

func lookupIP(domain string) ([]net.IP, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
