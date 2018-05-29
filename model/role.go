package model

type Role struct {
	Name string
}

type Roles []Role

func (r Roles) GetRoleAsSliceString() []string{
	s := make([]string, 0)
	for _, item := range r{
		s = append(s, item.Name)
	}

	return s
}