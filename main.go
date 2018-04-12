package main

const (
	MODE_USE     = "use"
	MODE_INIT    = "init"
	MODE_ROLE    = "role"
	MODE_MACHINE = "machine"
	MODE_EXE     = "exe"
)

var Mode = []string{
	MODE_USE,
	MODE_INIT,
	MODE_ROLE,
	MODE_MACHINE,
	MODE_EXE,
}

func main() {
	pargs := New()
	switch pargs.Mode {
	case MODE_INIT:
		doInit(pargs)
	case MODE_USE:
		doUse(pargs)
	}
}
