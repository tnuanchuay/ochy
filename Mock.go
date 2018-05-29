package main

func NewProjectProgramArgsMock(
	projectName string,
	args []string,
	mode string,
)ProjectProgramArgs{
	return ProjectProgramArgs{
		&ProgramArgs{
			ProjectName: projectName,
			Args:        args,
			Mode:        mode,
		},
		&Project{
			ProjectName: projectName,
		},
	}
}
