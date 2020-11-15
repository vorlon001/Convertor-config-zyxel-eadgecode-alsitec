package main

import (
    "fmt"
    "path/filepath"
    "runtime"
    "time"
)

func main() {
    getFileName(1)
    time.Sleep(time.Hour)
}

func getFileName(shift int) {
    go func() {
        _, file, line, ok := runtime.Caller(shift)
        if !ok {
            file = "???"
            line = 0
        } else {
            file = filepath.Base(file)
        }

        fmt.Printf("1)%s:%d\n", file, line)
    }()
    _, file, line, _ := runtime.Caller(shift)
   fmt.Printf("2)%s:%d\n", file, line)
}
