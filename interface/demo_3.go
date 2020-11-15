package main

import (
	"fmt"
)
func H1(z interface{}) {
	fmt.Println(z)
}
func H2(z *interface{}) {
	fmt.Println(*z)
}
func H3(z *interface{}) {
	f := (*z).(*string)
	fmt.Printf("H3) %#v\r",*f)
}
func toStringInterface(a string)  interface{} {
	return a
}
func toStringInterfacePTR(a *string)  interface{} {
	return a
}
func main() {
	H1("Hello, playground");
	h2 := toStringInterface("Hello, playground")
	H2(&h2);
	h30 := "Hello, playground"
	h3 := toStringInterfacePTR(&h30)
	H3(&h3);
	
	var a interface{}
	z := 4_555_555
	x := &z
	a  = &x
	b := &a
	c := &b
	d := &c	
	
	fmt.Printf("H4) %#v\r", d)
	fmt.Printf("H4) %#v\r", **(*(*(*d))).(**int) )	
}
