package main

import (
	"github.com/tspn/ochy/util"
	"fmt"
)

const (
	MODE_ROLE_ADD = "add"
	MODE_ROLE_REMOVE = "remove"
)

var SubOpsRole = []string{
	MODE_ROLE_ADD,
	MODE_ROLE_REMOVE,
}

type RoleProgramArgs struct{
	ProjectProgramArgs
	SubCommand string
	RoleName   string
}

func NewRoleProgramArgs(pargs ProjectProgramArgs)RoleProgramArgs{
	rpargs := RoleProgramArgs{
		SubCommand:  getSubOps(pargs),
		RoleName:    getRoleValue(pargs),
		ProjectProgramArgs: pargs,
	}

	return rpargs
}

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

	panic(fmt.Sprintf("RoleName %s not found", role))
}

func doRole(pargs ProjectProgramArgs){
	rpargs := NewRoleProgramArgs(pargs)
	switch rpargs.SubCommand {
	case MODE_ROLE_ADD:
		rpargs.AddRole(rpargs.RoleName)
	case MODE_ROLE_REMOVE:
		rpargs.RemoveRole(rpargs.RoleName)
	}
}
func getRoleValue(pargs ProjectProgramArgs) string{
	defer func(){
		if v := recover(); v != nil{
			panic("role command require role name")
		}
	}()
	return pargs.Args[3]
}
func getSubOps(pargs ProjectProgramArgs) string{
	for _, v := range SubOpsRole{
		if pargs.Args[2] == v{
			return v
		}
	}

	panic(fmt.Sprintf("unknow subcommand %s", pargs.Args[2]))
}