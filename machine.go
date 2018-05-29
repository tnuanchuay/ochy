package main

import (
	"net"
)

type Machine struct {
	Name string
	Ip   net.IPAddr
	Role []string
}

