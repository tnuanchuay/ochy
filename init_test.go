package main

import (
	"testing"
	"os"
)

func TestDoInitWhenGet2Params(t *testing.T){
	defer func(){
		if r := recover(); r == nil{
			t.Error("expect panic, actual not panic")
		}
	}()
	os.Args = []string{"ochy", "init", }
	pargs := NewProgramArgs()
	doInit(pargs)
}