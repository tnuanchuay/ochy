package main

import (
	"fmt"
	"io/ioutil"
	"github.com/tspn/ochy/util"
	"sync"
	"os"
)

type Project struct {
	ProjectName string
	Machines    []Machine
	Roles       []string
	roleMutex   sync.Mutex
}

type ProjectProgramArgs struct{
	ProgramArgs
	*Project
}

func NewProjectProgramArgs(args ProgramArgs)ProjectProgramArgs{
	ppargs := ProjectProgramArgs{
		ProgramArgs: args,
		Project: Load(args.ProjectName),
	}

	return ppargs
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
func (p Project) Delete() {
	os.Remove(p.Filename())
}
func (p Project) GetMachine(name string) *Machine{
	for _, m := range p.Machines{
		if m.Name == name{
			return &m
		}
	}

	return nil
}

func NewProject(projectname string) *Project {
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
