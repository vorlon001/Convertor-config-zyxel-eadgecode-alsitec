package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func print() {
	fmt.Println("output")
	for i := 0; i < 64<<10; i++ {
		fmt.Println()
	}
}

func main() {
	old := os.Stdout // keep backup of the real stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = w

	// print() // fails if called here "fatal error: all goroutines are asleep - deadlock"

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	
	print() // works fine here

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	// reading our temp stdout
	fmt.Println("previous output size:", len(out))
	//fmt.Print(out)
}
