package model

import (
	"net"
)

type Machine struct {
	Name string
	Ip   net.IPAddr
	Role Roles
}

func (m *Machine) AddRole(role Role){
	m.Role = append(m.Role, role)
}

