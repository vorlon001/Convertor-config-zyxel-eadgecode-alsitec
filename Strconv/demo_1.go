package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {

reg0 := [...]string {"a","b","c"}
fmt.Println(strings.Join(reg0[:],","))

reg := []string {"a","b","c"}
fmt.Println(strings.Join(reg,","))
reg1 := []int{2345,2345,2345}
fmt.Println(len(reg1))
reg2 := make([]string,0)
for _,v := range reg1{
	reg2 = append(reg2 , strconv.Itoa(v))
}

fmt.Println(strings.Join(reg2,","))
}
