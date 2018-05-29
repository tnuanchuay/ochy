package main

import (
	"fmt"
	"io/ioutil"

	"github.com/tspn/SliceUtils"
	"github.com/tspn/ochy/util"
)

func doUse(pargs *ProgramArgs) {
	projectname := getProjectFile(pargs)
	if util.FileExist(projectname) {
		err := ioutil.WriteFile(".ochy", []byte(projectname), 0666)
		if err != nil {
			fmt.Println(projectname)
			util.Exit(err)
		}
		fmt.Printf("Use project %s\n", projectname)
	} else {
		util.Exit("Project file doesn't exist")
	}
}

func GetCurrentConfig() string {
	b, err := ioutil.ReadFile(".ochy")
	if err != nil {
		return ""
	}

	return string(b)
}

func getProjectFile(pargs *ProgramArgs) string {
	su := sliceutils.String(pargs.Args)
	return su.Tail().(string)
}
