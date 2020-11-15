package main

import (
	"fmt"
	"time"
	"sort"
)

func main() {
	layout := "2006-08-01"
	t, e := time.Parse("2006-01-02",layout)
	t = t.AddDate(0, 0, -1)
	fmt.Printf("%v %v",t,e)
	
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	fmt.Printf("%v \n",t)
	
	
	f := 65
	r := map[int]*int{ 65: &f }
	fmt.Printf("1)%v \n",r[1])

	e1 := 65
	m1 := 65
	z1 := 0
	t1 := map[*int]*int{ &m1: &e1 }
	h1 := t1[&m1]
	if h1!=nil {
		fmt.Printf("2)%v \n",*h1)
	}
	h2 := t1[&z1]
	if h2!=nil {
		fmt.Printf("3)%v \n",*h2)
	}


	m := map[int]int{44: 323, 33: 3332, 22: 33325}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}

}
