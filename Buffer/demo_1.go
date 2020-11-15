package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello "))
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)

	b.Write([]byte("Hello "))
	b.Write([]byte("Hello "))
	

	fmt.Printf("%#v \n",b.Bytes())
		
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)
}
