package lib

import "net"


// get te IP address of a domain
func LookupIP(domain string) ([]net.IP, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
