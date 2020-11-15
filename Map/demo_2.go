package main

import (
    "fmt"
)

func main() {

	r2 := []map[int]int{{0:11, 1:21, 2:31}}
	for k,v := range r2 {
		for i := range v {
            		fmt.Println("r2",k,v[i])
    		}
	}

	type z []map[string]interface{};
	r3 := z{{"0":11, "1":21, "3":31}}
	for k,v := range r3 {
		for i := range v {
            		fmt.Println("r3",k,v[i])
    		}
	}

	type e map[string]map[string]interface{};
	r4 := e{"V":{"0":"sdfgsdf", "1":21, "3":31}}
	for k,v := range r4 {
		for i := range v {
	                fmt.Println("r4",k,v[i])
	        }
	}

	type e1 map[string]interface{};
	type e2 map[string]e1;
	r5 := e2{"V":{"0":"sdfgsdf", "1":21, "3":31}}
	for k,v := range r5 {
		for i := range v {
            		fmt.Println("r5",k,v[i])
    		}
	}

}
