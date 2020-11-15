package main

import (
	"fmt"
	"math"
)

type B = struct {
	x string
	y bool
}
type K = struct {
	a int
	b B
}
type geometry interface {
	area() float64
	perim() float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
	
func main() {



	type T1 = map[K]geometry 
	type T2 = map[K]interface{}	
	type T3 = map[interface{}]interface{}	
	type N1 = []map[K]geometry 
	type N2 = []map[interface{}]interface{}	
		
	var k = K{ a:4 , b:B{x:"test"}}
	
	c := circle{radius: 5}

	fmt.Println(k)	
	fmt.Println(c)
	fmt.Println(c.area())
	fmt.Println(c.perim())

	var m1 = T1{k: c}

	fmt.Printf("%T %3v \n",m1[k],m1[k])
	d1 := m1[k].(circle)
	fmt.Println(d1.area())
	fmt.Println(d1.perim())


	var m2 = T2{k: c}

	fmt.Printf("%T %3v \n",m2[k],m2[k])
	d2 := m2[k].(circle)
	fmt.Println(d2.area())
	fmt.Println(d2.perim())
	
	var m3 = T3{k: c}

	fmt.Printf("%T %3v \n",m3[k],m3[k])
	d := m3[k].(circle)
	fmt.Println(d.area())
	fmt.Println(d.perim())

	var t1 = N1{{k: c}}

	fmt.Println(t1)
	fmt.Println(t1[0][k])
	f1 := t1[0][k]
	fmt.Println(f1.area())
	fmt.Println(f1.perim())

	

	var t2 = N2{{k: c},{1:c},{1:2}}

	fmt.Println(t2)
	fmt.Println(t2[0][k])
	f2 := t2[0][k].(circle)
	fmt.Println(f2.area())
	fmt.Println(f2.perim())

	
	fmt.Println("Hello, playground")
}
