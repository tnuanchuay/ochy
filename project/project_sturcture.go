package project

import (
	"fmt"
	"io/ioutil"
	"github.com/tspn/ochy/util"
	"github.com/tspn/ochy/model"
)

type Project struct {
	ProjectName string
	Machines    []model.Machine
	Roles       model.Roles
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
func (p Project) GetAllRoleName() []string{
	s := make([]string, 0)
	for _, ss := range p.Roles{
		s = append(s, ss.Name)
	}

	return s
}
func (p Project) GetRole(role string) *model.Role{
	for _, r := range p.Roles{
		if r.Name == role{
			return &r
		}
	}

	return nil
}
func (p Project) GetMachine(name string) *model.Machine{
	for _, m := range p.Machines{
		if m.Name == name{
			return &m
		}
	}

	return nil
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
