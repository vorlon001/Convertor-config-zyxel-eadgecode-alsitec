package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	Name        string
	Designation string
	Department  string
	Salary      int
	Email       string
}

func appendItem(items *strings.Builder, item string) {
	if len(item) > 0 {
		if items.Len() > 0 {
			items.WriteByte(' ')
		}
		items.WriteString(item)
	}
}

func (e Employee) String() string {
	s := new(strings.Builder)
	appendItem(s, e.Name)
	appendItem(s, e.Designation)
	appendItem(s, e.Department)
	appendItem(s, e.Email)
	return s.String()
}

func main() {
	ee := Employee{
		Name:        "John Smith",
		Designation: "Manager",
		Department:  "Sales",
		Email:       "john.smith@example.com",
		Salary:      42000,
	}
	fmt.Println(ee)
}
