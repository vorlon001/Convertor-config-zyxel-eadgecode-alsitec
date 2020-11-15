package main

import (
	"fmt"
	"time"
//	"bytes"
	"io"
)

// The pinger prints a ping and waits for a pong
func pinger(pw *io.PipeWriter, pinger <-chan int, ponger chan<- int) {
	i := 0
	for {
		<-pinger
		ponger <- i
		fmt.Printf("ping: %d\n", i)
		go func() {
			s:= fmt.Sprintf("ping:%d",i)
			a,b := pw.Write( []byte(s)  )		
			fmt.Printf("SEND %s PING! SEND BYTE:%#v ERROR:%#v \n",s,a,b)
		}()
		
		time.Sleep(time.Second)
		i++
		
	}
}

// The ponger prints a pong and waits for a ping
func ponger(pr *io.PipeReader, pinger chan<- int, ponger <-chan int) {
	for {
		i := <-ponger
		fmt.Printf("pong: %d\n", i)
		go func() {
			buffer := make([]byte, 1000)
			count, _ := pr.Read(buffer)
			if count > 0 {
				fmt.Printf("from PING: %s\n", buffer[:count])
			}
		}()		
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {

	pr, pw := io.Pipe()
	fmt.Printf(" %#v %#v \n",pr,pw)
	ping := make(chan int)
	pong := make(chan int)
	go pinger(pw,ping, pong)
	go ponger(pr,ping, pong)
	ping <- 1
	time.Sleep(4*time.Second)
}
