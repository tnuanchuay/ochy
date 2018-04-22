package main

const (
	OPERATION_ADD		=	"add"
	OPERATION_DELETE	=	"del"
)

type RoleArgs struct{
	ProgramArgs		*ProgramArgs
	Operation		string
	MachineName		string
	Ip				string
}

func NewRoleArgs (args *ProgramArgs) *RoleArgs{
	ra := RoleArgs{}
	ra.Operation = getOperationFromProgramArgs(args)
	if ra.Operation == ""{
		panic("invalid operation")
	}

	ra.MachineName = args.Args[3]
	ra.Ip = args.Args[4]

	return &ra
}

func doRole(pargs *ProgramArgs){

}

func getOperationFromProgramArgs(args *ProgramArgs) string{
	for _, ops := range []string{OPERATION_ADD, OPERATION_DELETE}{
		if  args.Args[2] == ops{
			return ops
		}
	}

	return ""
}