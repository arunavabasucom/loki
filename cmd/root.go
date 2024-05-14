package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"loki.projects.arunavabasu.com/lib"
)

var rootCmd = &cobra.Command{
	Use:   "loki [domain]",
	Short: "Modern CLI for DNS lookups",
	Long:  `loki is a modern CLI for DNS lookups. This application allows you to perform DNS lookups for a given hostname.`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[38], 100*time.Millisecond)
		s.Suffix = " looking for ips"
		s.Start()
		time.Sleep(1 * time.Second) 
		domain := args[0]

		// getteing the IP address of the domain
		ips, err := lib.LookupIP(domain)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		s.Stop()
		fmt.Printf("IP addresses for %s:\n", domain)
		for _, ip := range ips {
			fmt.Println(ip)
		}
	},
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
