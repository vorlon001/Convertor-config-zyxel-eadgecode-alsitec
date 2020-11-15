package main

import (
	"fmt"
	"runtime"
)

func trace2() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	fmt.Printf("%#v\n",frames);
	frame, _ := frames.Next()
	fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}


// Original from https://stackoverflow.com/a/25927915/1490379
func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}

func main() {
	trace2() // Returns correct line number (27)
	trace() // Returns line 29 (should be line 28)
}
