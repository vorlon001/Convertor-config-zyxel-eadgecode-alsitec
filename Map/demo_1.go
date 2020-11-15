package main

import (
	"fmt"
	"sort"
)

func main() {
	d := map[int]int{0:2, 1:620, 2:8, 3:8, 4:0, 5:0, 6:0, 7:1, 8:0, 9:112, 10:1, 11:0, 12:39, 13:1, 14:0}
	for k,v := range d {
		fmt.Println(k,v)
	}
	
	var keys []int
	for k := range d {
        	keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
        	fmt.Println("Key:", k, "Value:", d[k])
	}
}
