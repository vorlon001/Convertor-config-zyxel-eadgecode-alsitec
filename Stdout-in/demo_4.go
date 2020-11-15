package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

/* Function to run the groutine to run for stdin read */
func read(r io.Reader) <-chan string {
    lines := make(chan string)
    go func() {
        defer close(lines)
        scan := bufio.NewScanner(r)
        for scan.Scan() {
            lines <- scan.Text()
        }
    }()
    return lines
}

func main() {
    mes := read(os.Stdin) //Reading from Stdin
    for anu := range mes {
        fmt.Println("Message to stdout")
        fmt.Println(anu) //Writing to Stdout
    }
}
