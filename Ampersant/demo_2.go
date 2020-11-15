package main

import "fmt"

func sum(s []int, c *chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("%#v\n",c)
	*c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	v := make(chan int)
	c := &v
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-*c, <-v // receive from c

	fmt.Println(x, y, x+y)
}
