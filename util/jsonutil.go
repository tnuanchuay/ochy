package util

import (
	"encoding/json"
	"fmt"
)

func StructToJson(i interface{}) string{
	s, err := json.Marshal(i)
	if err != nil{
		fmt.Println(err)
	}

	return string(s)
}

func JsonToStruct(j string, i interface{}){
	err := json.Unmarshal([]byte(j), i)
	if err != nil{
		fmt.Println(err)
	}
}