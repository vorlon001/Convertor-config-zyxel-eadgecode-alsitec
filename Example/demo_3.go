package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	i, err := strconv.ParseInt("1593690766", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm)
}
