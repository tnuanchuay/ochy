package main

import (
	"testing"
	)

func TestAddRole(t *testing.T){
	pj := NewProject("prod")
	pj.AddRole("hk-prod")
	if pj.Roles[0] != "hk-prod"{
		t.Errorf("expect %s, actual %s", "hk-prod", pj.Roles[0])
	}
}
func TestRemoveRole(t *testing.T){
	pj := NewProject("prod")
	pj.AddRole("hk-prod")

	if v := pj.RemoveRole("hk-prod"); v != true{
		t.Errorf("expect %s, actual %t", "true", v)
	}

	if len(pj.Roles) != 0{
		t.Errorf("expect %d, actual %d", 0, len(pj.Roles))
	}
}
func TestRemoveEmptyRole(t *testing.T){
	pj := NewProject("prod")
	defer func (){
		if recover() == nil{
			t.Errorf("Remove empty role should panic")
		}
	}()

	if v := pj.RemoveRole("hk-prod"); v != true{
		t.Errorf("expect %s, actual %t", "true", v)
	}

	if len(pj.Roles) != 0{
		t.Errorf("expect %d, actual %d", 0, len(pj.Roles))
	}
}