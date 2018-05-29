package main

import (
	"github.com/tspn/ochy/util"
	"fmt"
)

const (
	MODE_ROLE_ADD = "add"
	MODE_ROLE_REMOVE = "remove"
	MODE_ROLE_LIST = "list"
)

var SubOpsRole = []string{
	MODE_ROLE_ADD,
	MODE_ROLE_REMOVE,
	MODE_ROLE_LIST,
}

type RoleProgramArgs struct{
	ProjectProgramArgs
	SubCommand string
	RoleName   string
}

func NewRoleProgramArgs(pargs ProjectProgramArgs)RoleProgramArgs{
	rpargs := RoleProgramArgs{
		SubCommand:  getSubOps(pargs),
		ProjectProgramArgs: pargs,
	}

	if rpargs.SubCommand != MODE_ROLE_LIST{
		rpargs.RoleName = getRoleValue(pargs)
	}

	return rpargs
}

func (p *Project) AddRole(role string){
	p.roleMutex.Lock()
	for _, v := range p.Roles{
		if v == role{
			util.Exit(fmt.Sprintf("Role name %s already exists", role))
		}
	}
	p.Roles = append(p.Roles, role)
	fmt.Println(fmt.Sprintf("Role %s has been added", role))
	p.Save()
	p.roleMutex.Unlock()
}
func (p *Project) RemoveRole(role string){
	p.roleMutex.Lock()
	for i, v := range p.Roles{
		if v == role{
			p.Roles = util.RemoveStringSlice(p.Roles, i)
			p.Save()
			p.roleMutex.Unlock()
			fmt.Println(fmt.Sprintf("Role %s has been removed", role))
			return
		}
	}

	util.Exit(fmt.Sprintf("Role name %s not found", role))
}
func (p *Project) List(){
	p.roleMutex.Lock()
	fmt.Println("Roles")
	for _, r := range p.Roles{
		fmt.Printf("\t- %s\n", r)
	}
	p.roleMutex.Unlock()
}

func doRole(pargs ProjectProgramArgs){
	rpargs := NewRoleProgramArgs(pargs)
	switch rpargs.SubCommand {
	case MODE_ROLE_ADD:
		rpargs.AddRole(rpargs.RoleName)
	case MODE_ROLE_REMOVE:
		rpargs.RemoveRole(rpargs.RoleName)
	case MODE_ROLE_LIST:
		rpargs.List()
	}
}
func getRoleValue(pargs ProjectProgramArgs) string{
	defer func(){
		if v := recover(); v != nil{
			util.Exit("Role command require role name")
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

	util.Exit(fmt.Sprintf("unknow subcommand %s", pargs.Args[2]))
	return ""
}