package main

import "net"

type Project struct{
	Machines	[]Machine
	Roles		[]Role
}

type Role struct{
	Name	string
}

type Machine struct{
	Name	string
	Ip		net.IPAddr
	Role	[]Role
}
