package main

import (
	"fmt"
	"log"
	"runtime"
)

func HandleError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log where
		// the error happened, 0 = this function, we don't want that.
		_, fn, line, _ := runtime.Caller(1)		
		log.Printf("[error] %s:%d %v", fn, line, err)

		_, fn, line, _ = runtime.Caller(2)
		log.Printf("[error] %s:%d %v", fn, line, err)		
		
		_, fn, line, _ = runtime.Caller(3)
		log.Printf("[error] %s:%d %v", fn, line, err)		
		
		b = true
	}
	return
}

func FancyHandleError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		pc, fn, line, _ = runtime.Caller(2)
		log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
		b = true
	}
	return
}

func main() {
	if FancyHandleError(fmt.Errorf("it's the end of the world")) {
		log.Print("stuff")
	}
}
