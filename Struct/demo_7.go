package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type T struct {
	A int64
	B float64
}
func (t *T) print() {
	fmt.Printf("%[1]v %[1]T \n",(*t).B)
}

func main() {
	// Create a struct and write it.
	t := T{A: 0xEEFFEEFF, B: 3.14}
	t.print()
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, t)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.Bytes())

	// Read into an empty struct.
	t1 := T{}
	err = binary.Read(buf, binary.BigEndian, &t1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x %f", t1.A, t1.B)
	t1.print()
}
