package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
}

func (person *Person) GetName() string {
	return person.Name
}

func (person Person) GetName2() string {
	return person.Name
}

func main() {
	person := Person{Name: "John Doe"}
	t := reflect.TypeOf(&person)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%#v %#v \n",m.Name,m.Func)
	}
}
