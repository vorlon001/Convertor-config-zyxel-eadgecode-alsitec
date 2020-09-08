package main

import (
	"fmt"
)

type Node struct {
	Name interface{}
}
type NodePtr  struct {
}
func (c NodePtr) GetB(b interface{}) *Node {
	return b.(*Node)
}
func (c NodePtr) SetB(b interface{}, data interface{}) {
	c.GetB(b).Name = data
}
func (c NodePtr) NewB()  interface{} {
	return  &Node{Name:nil}
}

type NodeStorage struct {
	t *map[string]interface{}
}
func (t *NodeStorage) New() {
	tb := make(map[string]interface{})
	(*t).t = &tb
}
func (t *NodeStorage) GetTB(key string) interface{} {
	return (*t.t)[key]
}
func (t *NodeStorage) IsKeyTB(key string, n func () interface{}) interface{} { 
	if _,ok:=(*t.t)[key]; !ok {	
		(*t.t)[key] = n()
	}
	return (*t.t)[key]
}

func (t *NodeStorage) GetData(key string) interface{} {
	c := NodePtr{}
	f := c.GetB(t.GetTB(key))
	return f.Name
}

type Data struct {
	I int
	Name string
}
func main() {

	c := NodePtr{}
	
	key := "4523"
	tb := NodeStorage{}
	tb.New()
	b := tb.IsKeyTB(key,c.NewB)
	i := Data{I: 6_666_777, Name: "23 234 % sdfxgdfg "}
	c.SetB(b,&i)
	y := tb.GetData(key)
	o := y.(*Data)
	fmt.Printf(" %#[1]v \n", *o )
	fmt.Printf(" %#[1]v \n", tb )
}
