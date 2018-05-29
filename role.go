package main

import (
	"github.com/tspn/ochy/util"
	"fmt"
)

func (p *Project) AddRole(role string){
	p.roleMutex.Lock()
	p.Roles = append(p.Roles, role)
	p.roleMutex.Unlock()
}

func (p *Project) RemoveRole(role string) bool{
	p.roleMutex.Lock()
	for i, v := range p.Roles{
		if v == role{
			p.Roles = util.RemoveStringSlice(p.Roles, i)
			return true
		}
	}

	panic(fmt.Sprintf("Role %s not found", role))
}