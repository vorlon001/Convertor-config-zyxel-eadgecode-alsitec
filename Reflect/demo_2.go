package main

import "fmt"
import "reflect"

type A struct {
   Foo string
   F string
   B string
}

func (a *A) PrintFoo(){
    fmt.Println("Foo value is " + a.Foo)
}

func main() {
        a := &A{Foo: "afoo", F: "2345234", B: " 234 234 "}

        val := reflect.ValueOf(a).Elem()
	for i := 0; i < val.NumField(); i++ {
		value := val.Field(i)
		name := val.Type().Field(i).Name
		types := val.Type().Field(i).Type
		fmt.Printf("%#v %#v %#v  %v\n",i, value, name, types )
	}
}
