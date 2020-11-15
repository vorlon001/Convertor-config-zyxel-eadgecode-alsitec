package main

import ( 
	"fmt"
	"runtime"
	"log"
)

func main() {
    f()
    fmt.Println("Returned normally from f.")
}


func panicrecover() {
        if r := recover(); r != nil {
		fmt.Println(">>>Recovered in f", r)
		log.Printf("Internal error: %v", r)
		buf := make([]byte, 1<<16)
		log.Fatal("2435234,52435243",234,"2345")
		stackSize := runtime.Stack(buf, true)
		log.Printf(">>>%s\n", string(buf[0:stackSize]))
	}
    }


func f() {
    defer panicrecover()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
