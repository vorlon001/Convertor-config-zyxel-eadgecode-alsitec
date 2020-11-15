package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	s := `{"text":"I'm a text.","number":1234,"floats":[1.1,2.2,3.3],"d":{"text":"I'm a text.","number":1234},
		"innermap":{"foo":1,"bar":2}}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		panic(err)
	}

	if v, ok := data["d"]; ok {
		fmt.Println("KEY 'd1' found", v)
		fmt.Println("text =", v)
		h,err := v.(map[string]interface{})
		if !err {
			panic("inner map is not a map!")
		}
		fmt.Println("text =", h["text1"])
		keys := reflect.ValueOf(data).MapKeys()
		fmt.Println("text =", keys)

	} else {
		fmt.Println("KEY 'd' not found")
	}

	if v, ok := data["d1"]; ok {
		fmt.Println("KEY 'd1' found", v)
	} else {
		fmt.Println("KEY 'd1' not found")
	}

	fmt.Println("text =", data["text"])
	fmt.Println("number =", data["number"])
	fmt.Println("floats =", data["floats"])
	fmt.Println("innermap =", data["innermap"])

	innermap, ok := data["innermap"].(map[string]interface{})
	if !ok {
		panic("inner map is not a map!")
	}
	fmt.Println("innermap.foo =", innermap["foo"])
	fmt.Println("innermap.bar =", innermap["bar"])

	fmt.Println("The whole map:", data)
}
