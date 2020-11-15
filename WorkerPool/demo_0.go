package main

import (
	"fmt"
	"time"
//	"os"
	"sync"
)

func pinger(pinger chan<- *interface{}, stop chan<- bool, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	j :=0
	for {
		if j>30 {
			fmt.Println("pinger - init send stop and close")
			for { 
				if len(pinger )==0 {
					break
				}
				time.Sleep(1000 * time.Millisecond)
			}
			stop <- true
			close(done)
			fmt.Println("pinger - stop is done")
			return
		}
		fmt.Println("ping ",j)
		time.Sleep(100 * time.Millisecond)
		var s interface{}
		s = j
		pinger <- &s
		j++		
	}
}

func ponger(pinger <-chan *interface{}, stop <-chan bool, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
			case val := <-pinger:
				fmt.Println("pong ",*val)
				time.Sleep(300 * time.Millisecond)
			case <-stop:
				fmt.Println("ponger - stop is done")
				close(done)
				return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ping := make(chan *interface{},30)
	stop := make(chan bool)
	doneping := make(chan bool)
	donestop := make(chan bool)
	
	wg.Add(1)
	go pinger(ping, stop, doneping, &wg)
	wg.Add(1)
	go ponger(ping, stop, donestop, &wg)
	wg.Wait()
	close(ping )
	close(stop )
	<-doneping 
	<-donestop 
}

