package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	r := ""
	for i := 0; i < len(ip); i++ {
		var concatenated string
		if i == 3 {
			concatenated = fmt.Sprint(ip[i])
		} else {
			concatenated = fmt.Sprint(ip[i], ".")
		}
		r += concatenated
	}
	return r
}

func stringex() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
