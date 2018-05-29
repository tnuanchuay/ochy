package main

import (
	"fmt"
	"github.com/tspn/ochy/util"
)

func doInit(args *ProgramArgs){
	if len(args.Args) != 3 {
		util.Exit(fmt.Sprintf("Require 2 Arguments.\nfor example: ochy init [project name]",))
	}

	pn := args.Args[2]
	ps := NewProject(pn)
	ps.Save()
	fmt.Printf("Initialize project %s\n", filename(ps.ProjectName))
}