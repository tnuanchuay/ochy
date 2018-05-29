package main

import (
	"os"

	"github.com/tspn/SliceUtils"
)

type ProgramArgs struct {
	Mode        string
	ProjectName string
	Args        []string
}

func NewProgramArgs() *ProgramArgs {
	return newProgramArgs(os.Args)
}
func newProgramArgs(s []string) *ProgramArgs {
	pargs := new(ProgramArgs)
	pargs.Mode = getMode(s)
	if pargs.Mode != "" {
		pargs.Args = os.Args
		pargs.ProjectName = GetCurrentConfig()
		return pargs
	} else {
		return nil
	}
}
func getMode(s []string) string {
	mode := sliceutils.String(Mode).
		Filter(func(elem interface{}) bool {
			if len(s) > 1 {
				return elem.(string) == s[1]
			} else {
				return false
			}
		})

	if len(mode) != 1 {
		return ""
	} else {
		return mode.Head().(string)
	}
}
func hasMode(s []string) bool {
	return sliceutils.String(Mode).Filter(func(elem interface{}) bool {
		if len(s) > 1 {
			return elem.(string) == s[1]
		} else {
			return false
		}
	}).Len() > 0
}
