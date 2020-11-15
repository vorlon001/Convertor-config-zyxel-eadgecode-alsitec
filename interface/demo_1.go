package main

import (
	"fmt"
)
type Obj1 struct {
	Z string
}
func (a *Obj1) Show() {
	fmt.Println("A: Not Yet Set")
}

type Obj2 struct {
	Z string
}
func (d *Obj2) Draw() {
	fmt.Println("D: Not Yet Set")
}


type Obj3 struct {
	Obj1
	F string
}
func (b *Obj3) Show() {
	fmt.Printf("B: Show: %v - %v \n", b.Z, b.F)
}

func (b *Obj3) Draw() {
	fmt.Printf("B: Draw(): %v - %v \n", b.Z, b.F)
}

type Shape_Show interface {
    Show()
}

type Shape_Draw interface {
    Draw()
}

func Show(s Shape_Show ){
	s.Show();
}

func Draw(s Shape_Draw ){
	s.Draw();
}

func main() {
	fmt.Println("Hello, playground")
	
	a := Obj1{Z:"23452345 2 345$"}
	d := Obj2{Z:"23452345 2 345$"}
	b := Obj3{ Obj1: Obj1{Z:"23452345 2 345$"}, F: "sdgfasdfasdfasd" }
	
	a.Show()
	b.Show()

	Show(&a)
	Show(&b)

	Draw(&d)
	Draw(&b)
}
