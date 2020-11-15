package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	FirstName string `tag_name:"many2many" json:"locale_blog_tags1"`
	LastName  string `tag_name:"tag 2" json:"locale_blog_tags2"`
	Age       int    `tag_name:"tag 3" json:"locale_blog_tags3"`
}

type Foo2 struct {
	Foo Foo 
}

func (f *Foo) reflect() {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s,%s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name") , tag.Get("json"))
	}
} 

func (f *Foo2) reflect() {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s,%s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name") , tag.Get("json"))
	}
}
func main() {
	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}

	f.reflect()

}
