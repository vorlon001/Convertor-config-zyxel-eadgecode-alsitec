package main

import (
	"fmt"
	"net"
)


func main() {

	type net_obj struct { Network string "yaml:\"network\""; Token string "yaml:\"token\"" };
	
	network := []net_obj{{ Network: "10.228.14.0/24", Token: "saer"},{ Network: "10.228.15.0/24", Token: "saer"},{ Network: "0.0.0.0/0", Token: "saer"}}
	clientip:= "10.228.14.117"

	verify := func (network []net_obj, clientip string ) string  {
		ip := net.ParseIP(clientip)
		for _, v:= range network {
			_, subnet, _ := net.ParseCIDR(v.Network )
			if subnet.Contains(ip) {
				return v.Token
			} 
		}
		return "PUBLIC"
	};
	
	fmt.Println( verify(network , clientip) )
}
