package util

import (
	"fmt"
	//"os"
)

func Exit(err interface{}){
	fmt.Println(err)
	//os.Exit(1)
	panic(err)
}
