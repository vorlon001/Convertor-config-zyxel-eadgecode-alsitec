package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fmt.Println("Hello, playground") // this gets captured

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	fmt.Printf("Captured: %s", out)
}
