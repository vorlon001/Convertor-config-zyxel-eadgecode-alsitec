package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		buf1    bytes.Buffer
		logger1 = log.New(&buf1, "INFO: ", log.Lshortfile)

		buf2    bytes.Buffer
		logger2 = log.New(&buf2, "INFO: ", log.Lshortfile)

		infof1 = func(info string) {
			logger1.Println(info)
		}
		infof2 = func(info string) {
			logger2.Println(2, info)
		}
	)

	multi := io.MultiWriter(logger1.Writer(), logger2.Writer(), os.Stdout)

	log.SetOutput(multi)
	log.Println("message 1")

	infof1("Hello world")
	infof2("Hello world")

	fmt.Printf("---------")
	fmt.Printf("log1> %v\n", &buf1)
	fmt.Printf("---------")
	fmt.Printf("log2> %v\n", &buf2)
	fmt.Printf("---------")
}
