package project

import (
	"net"
	"github.com/tspn/ochy/util"
	"io/ioutil"
	"fmt"
)

type Project struct{
	ProjectName	string
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

func (p Project) ToJson() string{
	return util.StructToJson(p)
}
func (p Project) Filename() string{
	return filename(p.ProjectName)
}
func (p Project) Save(){
	err := ioutil.WriteFile(p.Filename(), []byte(p.ToJson()), 0666)
	if err != nil{
		panic("could not initial project file")
	}
}

func New(projectname string) *Project{
	return &Project{
		ProjectName:projectname,
	}
}
func Load(projectname string) *Project{
	b, err := ioutil.ReadFile(filename(projectname))
	if err != nil{
		panic(err)
	}
	var p Project
	util.JsonToStruct(string(b), &p)

	return &p
}
func filename(projectname string) string{
	return fmt.Sprintf("%s.config.json", projectname)
}