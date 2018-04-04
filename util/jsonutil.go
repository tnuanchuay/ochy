package util

import "encoding/json"

func StructToJson(i interface{}) string{
	s, err := json.Marshal(i)
	if err != nil{
		panic(err)
	}

	return string(s)
}

func JsonToStruct(j string, i interface{}){
	err := json.Unmarshal([]byte(j), i)
	if err != nil{
		panic(err)
	}
}