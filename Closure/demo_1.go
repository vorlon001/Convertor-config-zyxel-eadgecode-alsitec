package main

import (
	"fmt"
)

func closure1(s string) func(string){
    text := "Outer " + s
    anon := func(a string){
        fmt.Println(text,a)
    }
    return anon
}
func closure(s string) {
    text := "Outer " + s
    anon := func(){
        fmt.Println(text)
    }
    anon()
}
 
func main(){
    closure("space")
    i := closure1("space")
    i("2134");
    i("2dfgdsa");
}
