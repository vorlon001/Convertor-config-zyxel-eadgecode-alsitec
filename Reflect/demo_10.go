package main

import (
	"fmt"
	"reflect"
)

func main() {

	r4 := reflect.TypeOf(map[int]interface{}{1:"hi", 2:42, 3: func() {}})
	fmt.Println(r4)
	z4 := []interface{}{"hi", 42, func() {}}
	switch v := reflect.ValueOf(z4); v.Kind() {
		case reflect.Array:
			fmt.Println("Array",z4)
		case reflect.Map:
			fmt.Println("Map",z4)
		case reflect.Slice:
			fmt.Println("Slice",z4)

		default:
			fmt.Printf("unhandled <>kind %s", v.Kind())
	}


	r := reflect.TypeOf([]interface{}{"hi", 42, func() {}})
	fmt.Println(r)
	z := []interface{}{"hi", 42, func() {}}
	switch v := reflect.ValueOf(z); v.Kind() {
		case reflect.Array:
			fmt.Println("Array",z)
		case reflect.Map:
			fmt.Println("Map",z)
		case reflect.Slice:
			fmt.Println("Slice",z)

		default:
			fmt.Printf("unhandled <>kind %s", v.Kind())
	}

	z1:= map[string]int{
	    "rsc": 3711,
	    "r":   2138,
	    "gri": 1908,
	    "adg": 912,
	}
	z2 := map[int]int{1: 4,2: 4,3: 4,4: 4}
	z3 := [5]float64{ 98, 93, 77, 82, 83 }

	switch v := reflect.ValueOf(z1); v.Kind() {
		case reflect.Array:
			fmt.Println("Array",z1)
		case reflect.Map:
			fmt.Println("Map",z1)
		case reflect.Slice:
			fmt.Println("Slice",z1)
		default:
			fmt.Printf("unhandled <>kind %s\n", v.Kind())
	}

	switch v := reflect.ValueOf(z2); v.Kind() {
		case reflect.Array:
			fmt.Println("Array",z2)
		case reflect.Map:
			fmt.Println("Map",z2)
		case reflect.Slice:
			fmt.Println("Slice",z2)
		default:
			fmt.Printf("unhandled <>kind %s\n", v.Kind())
	}


	switch v := reflect.ValueOf(z3); v.Kind() {
		case reflect.Array:
			fmt.Println("Array",z3)
		case reflect.Map:
			fmt.Println("Map",z3)
		case reflect.Slice:
			fmt.Println("Slice",z3)
		default:
			fmt.Printf("unhandled <>kind %s\n", v.Kind())
	}

	
	for _, v := range []interface{}{"hi", 42, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}

}
