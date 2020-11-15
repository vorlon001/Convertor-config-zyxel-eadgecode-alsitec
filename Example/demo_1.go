package main

import (
	"fmt"
)


func PrintMemUsage() {

	bToMb := func (b uint64) uint64 {
		return b / 1024 / 1024
	}
	
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func main() {
	fmt.Println("Hello, playground")
}
