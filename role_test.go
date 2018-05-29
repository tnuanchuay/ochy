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

func TestGetSubOps(t *testing.T){
	ppargs := GiveProjectProgramArgsMockForRole([]string{ "ochy", "role", "add", "hk-prod"})

	if v := getSubOps(ppargs); v != MODE_ROLE_ADD{
		t.Errorf("expect %s, actual %s", MODE_ROLE_ADD, v)
	}
}
func TestGetRoleValue(t *testing.T){
	ppargs := GiveProjectProgramArgsMockForRole([]string{ "ochy", "role", "add", "hk-prod"})

	if v := getRoleValue(ppargs); v != "hk-prod"{
		t.Errorf("expect %s, actual %s", "hk-prod", v)
	}
}
func TestGetRoleValueNoRoleName(t *testing.T){
	ppargs := GiveProjectProgramArgsMockForRole([]string{ "ochy", "role", "add",})

	defer func() {
		if recover() == nil{
			t.Errorf("No Rolename should panic")
		}
	}()

	getRoleValue(ppargs)
}
func TestDoRoleAdd(t *testing.T){
	ppargs := GiveProjectProgramArgsMockForRole([]string{ "ochy", "role", "add", "hk-prod"})

	doRole(ppargs)
	if ppargs.Roles[0] != "hk-prod"{
		t.Errorf("expect %s, actual %s", "hk-prod", ppargs.Roles[0])
	}
}
func TestDoRoleRemove(t *testing.T){
	ppargsRemove := GiveProjectProgramArgsMockForRole([]string{ "ochy", "role", "remove", "hk-prod"})
	ppargsRemove.Roles = append(ppargsRemove.Roles, "hk-prod")

	doRole(ppargsRemove)

	if len(ppargsRemove.Roles) != 0{
		t.Errorf("expect %d, actual %d", 0, len(ppargsRemove.Roles))
	}
}

func GiveProjectProgramArgsMockForRole(args []string) ProjectProgramArgs{
	return NewProjectProgramArgsMock(
		"prod",
		args,
		MODE_ROLE,
	)
}

