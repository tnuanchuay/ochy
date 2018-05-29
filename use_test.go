package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/tspn/ochy/util"
)

func TestDoUseConfigDoesntExist(t *testing.T) {
	os.Args = []string{"ochy", "use", "prod.config.json"}
	defer func() {
		if v := recover(); v == nil {
			t.Errorf("expect panic because config file doesn't exist")
		}
	}()
	pargs := NewProgramArgs()
	doUse(pargs)
}

func TestDoUse(t *testing.T) {
	os.Args = []string{"ochy", "use", "prod.config.json"}
	ioutil.WriteFile("prod.config.json", []byte("{}"), 0666)
	defer os.Remove("prod.config.json")
	defer os.Remove("hk.config.json")
	defer os.Remove(".ochy")

	pargs := NewProgramArgs()
	doUse(pargs)
	if !util.FileExist(".ochy") {
		t.Errorf("expect .ochy file is created")
	}

	b, _ := ioutil.ReadFile(".ochy")
	if !reflect.DeepEqual(b, []byte("prod.config.json")) {
		t.Errorf("expect %s, actual %s", "prod.config.json", string(b))
	}

	os.Args = []string{"ochy", "use", "hk.config.json"}
	ioutil.WriteFile("hk.config.json", []byte("{}"), 0666)
	pargs = NewProgramArgs()
	doUse(pargs)
	b, _ = ioutil.ReadFile(".ochy")
	if !reflect.DeepEqual(b, []byte("hk.config.json")) {
		t.Errorf("expect %s, actual %s", "hk.config.json", string(b))
	}
}

func TestGetProjectFile(t *testing.T) {
	os.Args = []string{"ochy", "use", "./filename.config.json"}
	pargs := NewProgramArgs()
	if v := getProjectFile(pargs); v != os.Args[2] {
		t.Error(fmt.Sprintf("expect %s, actual %s", os.Args[2], v))
	}
}
