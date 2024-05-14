package lib

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <hostname>\n", os.Args[0])
		os.Exit(1)
	}

	hostname := os.Args[1]

	// Perform DNS lookup
	ips, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Print results
	fmt.Println("IP addresses for", hostname+":")
	for _, ip := range ips {
		fmt.Println(ip)
	}
}





