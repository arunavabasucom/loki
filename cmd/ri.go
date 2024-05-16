package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var riCmd = &cobra.Command{
	Use:   "ri [ip]",
	Short: "Reverse IP lookup",
	Long:  `Reverse IP lookup is a technique that allows you to find all the domains hosted on a given IP address.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
		s.Suffix = " looking for ips"
		s.Start()
		time.Sleep(1 * time.Second)
		ip := args[0]
		domains, err := net.LookupAddr(ip)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		s.Stop()

		fmt.Printf("Domain(s) associated with IP %s:\n", ip)
		for _, domain := range domains {
			fmt.Println(domain)
		}
	},
}

func init() {
	rootCmd.AddCommand(riCmd)
}
