package main

import (
	"fmt"
)

type VLAN struct {
			Name string `json:"name"`
			Tag  int    `json:"tag"`
		}

type PORT struct {
			Description string        `json:"description"`
			NativeVlan  int           `json:"native_vlan"`
			Status      string        `json:"status"`
			Tagged      []interface{} `json:"tagged"`
			Untagged    []int         `json:"untagged"`
		}

type SWITCH struct {
	DefaultGateway string        `json:"default_gateway"`
	DhcpSnooping   []int         `json:"dhcp_snooping"`
	Hostname       string        `json:"hostname"`
	IgmpSnooping   []int         `json:"igmp_snooping"`
	LoggingServer  []interface{} `json:"logging_server"`
	MngIntVlan     int           `json:"mng_int_vlan"`
	MngIP          string        `json:"mng_ip"`
	MngMask        string        `json:"mng_mask"`
	Ports          map[string]PORT `json:"pors"`
	SntpServer []string `json:"sntp_server"`
	Vlans      []VLAN `json:"vlans"`
}

func main() {
	fmt.Println("Hello, playground")
}
