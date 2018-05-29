package util

func RemoveStringSlice(slice []string, at int) []string{
	slice[len(slice)-1], slice[at] = slice[at], slice[len(slice)-1]
	return slice[:len(slice)-1]
}