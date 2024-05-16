package lib

import "net"


// LookupIP returns the IP addresses of a given domain
func LookupIP(domain string) ([]net.IP, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
