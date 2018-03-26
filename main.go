package main

const (
	MODE_INIT		= 		"init"
	MODE_CREATE		=		"create"
	MODE_REMOVE		=		"remove"
	MODE_ADD		=		"add"
)

var Mode = []string{
	MODE_INIT,
	MODE_CREATE,
	MODE_REMOVE,
	MODE_ADD,
}

func main(){
	pargs := New()
	switch pargs.Mode{
	case MODE_INIT:
		doInit(pargs)
	}
}