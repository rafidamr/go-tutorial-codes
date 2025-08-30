package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (addr IPAddr) String() string {
	var str string
	for _, b := range addr {
		str += fmt.Sprintf("%d", b) + "."
	}
	return fmt.Sprint(str[:len(str)-1])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Println(name, ip)
	}
}
