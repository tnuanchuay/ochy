package project

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/tspn/ochy/role"
	"github.com/tspn/ochy/util"
)

type Project struct {
	ProjectName string
	Machines    []Machine
	Roles       []role.Role
}

type Machine struct {
	Name string
	Ip   net.IPAddr
	Role []role.Role
}

func (p Project) ToJson() string {
	return util.StructToJson(p)
}
func (p Project) Filename() string {
	return filename(p.ProjectName)
}
func (p Project) Save() {
	err := ioutil.WriteFile(p.Filename(), []byte(p.ToJson()), 0666)
	if err != nil {
		panic("could not initial project file")
	}
}

func New(projectname string) *Project {
	return &Project{
		ProjectName: projectname,
	}
}
func Load(projectname string) *Project {
	b, err := ioutil.ReadFile(filename(projectname))
	if err != nil {
		panic(err)
	}
	var p Project
	util.JsonToStruct(string(b), &p)

	return &p
}

func filename(projectname string) string {
	return fmt.Sprintf("%s.config.json", projectname)
}
