package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
	"log"
)

func main() {
	fmt.Println("Hello, playground")
        buf := new(bytes.Buffer)
        w := io.MultiWriter(buf, os.Stderr)
        log.SetOutput(w)
	log.SetFlags(log.Llongfile)
	log.Println("hello, logfile")
	fmt.Println(buf )
}
