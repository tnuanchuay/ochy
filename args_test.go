package main

import (
	"testing"
	"fmt"
)

func Test_new(t *testing.T){
	s := []string{"ochy", "create", "role", "hk-prod"}
	pargs := _new(s)

	if pargs.Mode != MODE_CREATE {
		t.Error(fmt.Sprintf("expect %s, actual %s", MODE_CREATE, pargs.Mode))
	}

	s = []string{"ochy", "damn", "role", "hk-prod"}
	pargs = _new(s)
	if pargs != nil{
		t.Error(fmt.Sprintf("expect %s, actual %s", "nil", pargs))
	}

	s = []string{}
	pargs = _new(s)
	if pargs != nil{
		t.Error(fmt.Sprintf("expect %s, actual %s", "nil", pargs))
	}
}
func TestGetMode(t *testing.T){
	if r := getMode([]string{"ochy", "init"}); r != MODE_INIT{
		t.Error(fmt.Sprintf("expect %s, actual %s", MODE_INIT, r))
	}

	if r := getMode([]string{"ochy", "create", "role", "hk-production"}); r != MODE_CREATE{
		t.Error(fmt.Sprintf("expect %s, actual %s", MODE_CREATE, r))
	}

	if r := getMode([]string{"ochy", "mother", "fucker"}); r != ""{
		t.Error(fmt.Sprintf("expect %s, actual %s", "", r))
	}

	if r := getMode([]string{"ochy"}); r != ""{
		t.Error(fmt.Sprintf("expect %s, actual %s", "", r))
	}

	if r := getMode([]string{}); r != ""{
		t.Error(fmt.Sprintf("expect %s, actual %s", "", r))
	}
}
func TestIsMode(t *testing.T){
	if r := hasMode([]string{"ochy", "init"}); r != true{
		t.Error(fmt.Sprintf("expect %v, actual %v", true, r))
	}

	if r := hasMode([]string{"ochy", "create", "role", "hk-production"}); r != true{
		t.Error(fmt.Sprintf("expect %v, actual %v", true, r))
	}

	if r := hasMode([]string{"ochy", "mother", "fucker"}); r != false{
		t.Error(fmt.Sprintf("expect %v, actual %v", false, r))
	}

	if r := hasMode([]string{"ochy"}); r != false{
		t.Error(fmt.Sprintf("expect %v, actual %v", false, r))
	}

	if r := hasMode([]string{}); r != false{
		t.Error(fmt.Sprintf("expect %v, actual %v", false, r))
	}
}