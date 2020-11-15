package main

import (
	"fmt"
	"net"
	"log"
	"strings"
	"strconv"
	"bytes"
	"errors"
)

func isPrivate(ipv4 string) bool { 

	ip := net.ParseIP(ipv4)

	private := []string{
		`0.0.0.0/8`,
		`10.0.0.0/8`,
		`127.0.0.0/8`,
		`169.254.0.0/16`,
		`172.16.0.0/12`,
		`192.0.0.0/29`,
		`192.0.0.170/31`,
		`192.0.2.0/24`,
		`192.168.0.0/16`,
		`198.18.0.0/15`,
		`198.51.100.0/24`,
		`203.0.113.0/24`,
		`240.0.0.0/4`,
		`255.255.255.255/32`,
	}

	for _,v := range private {
		_, ipv4Net, err := net.ParseCIDR(v)
		if err != nil {
			log.Fatal(err)
		}
		if ipv4Net.Contains(ip) {
     	 		return true
		}
	}
	return false
}

func isPublic(ipv4 string) bool {
	if isPrivate(ipv4){
		return false
	}
	return true
}

func isIPv6(isipv6 string) (string,string,error) {
	ipv6bytes := []byte( isipv6 )
	maping := [][]string{ []string{"A","a"},[]string{"B","b"},[]string{"C","c"},[]string{"D","d"},[]string{"E","e"},[]string{"F","f"} }	
	
	ipv6maskint,err := strconv.ParseInt( string(ipv6bytes [0:4])  , 16, 64)
	if err!=nil {
		return "","",err
	}
	ipv6mask := strconv.FormatInt(ipv6maskint, 10)
	
	if len(ipv6bytes )==36 {
		ipv6bblock := make([]string,4)
		for k,v:= range []int{4,8,12,16} {
			block := ipv6bytes [v:v+4]
			for _,w := range maping {
				block = bytes.ReplaceAll(block , []byte(w[0]), []byte(w[1]))
			}
			ipv6bblock [k] = string(block)
		}
		return strings.Join(ipv6bblock,":"),ipv6mask , nil
	} 
	return "", "", errors.New("Error Parsing Point 100001")
}


func main() {
	fmt.Println( isPrivate("12.168.2.1") )
	fmt.Println( isPublic("12.168.2.1") )
	e := "00402A01054040026F2D0000000000000000"
	
	a,b,c := isIPv6(e)
	fmt.Printf("%v %v %v \n", a,b,c)
}
