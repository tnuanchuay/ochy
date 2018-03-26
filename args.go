package main

import (
	"github.com/tspn/SliceUtils"
	"os"
)

type ProgramArgs struct{
	Mode	string
	Args	[]string
}

func New() *ProgramArgs{
	return _new(os.Args)
}
func _new(s []string) *ProgramArgs{
	pargs := new(ProgramArgs)
	pargs.Mode = getMode(s)
	if pargs.Mode != ""{
		pargs.Args = os.Args
		return pargs
	}else{
		return nil
	}
}
func getMode(s []string) string{
	mode := sliceutils.String(Mode).
		Filter(func(elem interface{}) bool{
		if len(s) > 1{
			return elem.(string) == s[1]
		}else{
			return false
		}
	})

	if len(mode) != 1{
		return ""
	}else{
		return mode.Head().(string)
	}
}
func hasMode(s []string) bool{
	return sliceutils.String(Mode).Filter(func(elem interface{}) bool {
		if len(s) > 1{
			return elem.(string) == s[1]
		}else{
			return false
		}
	}).Len() > 0
}