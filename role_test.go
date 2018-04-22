package main

import "testing"

func TestNewRoleArgs(t *testing.T) {
	args := ProgramArgs{
		Mode: MODE_ROLE,
		Args:[]string{"ochy", "role", "add", "hk-prod", "127.0.0.1"},
		ConfigName:".ochy",
	}

	ra := NewRoleArgs(&args)
	if ra.Operation != OPERATION_ADD {
		t.Errorf("expect %s, actual, %s", OPERATION_ADD, ra.Operation)
	}

	if ra.MachineName != "hk-prod" {
		t.Errorf("expect %s, actual %s", "hk-prod", ra.MachineName)
	}

	if ra.Ip != "127.0.0.1" {
		t.Errorf("expect %s, actual %s", "127.0.0.1", ra.Ip)
	}
}