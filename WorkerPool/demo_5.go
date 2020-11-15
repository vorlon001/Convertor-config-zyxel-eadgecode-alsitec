package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

/*********************************************************************************************************************/
type HandlerError struct {
	Err         string
	Dump        string
	WorkerError *workerError
}

func (e *HandlerError) Error() string {
	return fmt.Sprintf("Err:%[1]s", e.Err)
}

/*********************************************************************************************************************/
type workerError struct {
	Id   int
	Err  string
	Dump string
}

func (e *workerError) Error() string {
	return fmt.Sprintf("workerError in worker Id:%[1]d Err:%[2]s", e.Id, e.Err)
}

/*********************************************************************************************************************/
func panicRecoverWorker(i chan<- *workerError, id int, quit chan bool, isDone *[]bool) {
	if r := recover(); r != nil {
		fmt.Printf("Run panicRecoverWorker(), Internal error: %v", r)
		buf := make([]byte, 1<<16)
		stackSize := runtime.Stack(buf, true)
		i <- &workerError{
			Id:   id,
			Dump: string(buf[0:stackSize]),
			Err:  fmt.Sprintf("Internal error: %v", r),
		}
		(*isDone)[id] = true
	}
}

/*********************************************************************************************************************/

func work(id int, max_sleep int) int {
	var n int = rand.Intn(max_sleep)
	fmt.Printf("run id:%v worker()\tsleep: %v msec.\n", id, n)
	/*
	a := 232
	b := 0
	if n < 1000 {
		_ = a / b
	}
	*/
	time.Sleep(time.Duration(n) * time.Millisecond)
	fmt.Printf("run id:%v worker()\tDone\n", id)
	return n
}

/*********************************************************************************************************************/

type POOL struct {
	maxPool int
	Pool    []int
	isDone  []bool
	Quit    []chan bool
	Done    chan bool
	Result  chan int
	Error   chan *workerError
}

func NewPool(maxPool int) *POOL {
	p := POOL{maxPool: 5,
		Pool:   make([]int, maxPool),
		Quit:   make([]chan bool, maxPool),
		Done:   make(chan bool, maxPool*20),
		isDone: make([]bool, maxPool),
		Result: make(chan int, maxPool*20),
	}
	for i := range p.Pool {
		p.Quit[i] = make(chan bool)
		p.isDone[i] = false

	}

	p.Error = make(chan *workerError, maxPool*10)

	return &p
}

func (p *POOL) Run(work func(int, int) int) {

	for i := range p.Pool {
		fmt.Printf("run id %v %#v\n", i, p.Quit[i])
		go p.Worker(i, (*p).Quit[i], p.Error, p.Result, work)
	}
	go p.Watcher(p.Done, p.Quit, p.Result)
	<-p.Done

}

func (p *POOL) Worker(id int, quit chan bool, chError chan<- *workerError, result chan<- int, worker func(int, int) int) {
	defer panicRecoverWorker(chError, id, quit, &p.isDone)

	i := 0
	for {
		select {
		case <-quit:
			fmt.Printf("worker() id:%v is resived command QUIT!\n", id)
			return
		default:
			fmt.Printf("\tworker() id:%v init: %v \n", id, i)
			var n int = worker(id, 5000)
			i++
			fmt.Printf("worker() id:%v iteration: %v \n", id, i)
			result <- n
		}
	}
}

func (p *POOL) Watcher(done chan<- bool, quit []chan bool, result <-chan int) {
	quitErrorHandler := make(chan bool)
	go func(quitErrorHandler <-chan bool) {
		fmt.Printf("init ERROR HANDLER WORKER\n")
		for {
			select {
			case <-quitErrorHandler:
				return
			case e := <-p.Error:
				fmt.Printf("ERROR in WORKER: \t%[1]v\n\t%#[1]v\n", e)
			}
		}
	}(quitErrorHandler)

	r := 0
	for {
		i := 5000
		time.Sleep(time.Duration(i) * time.Millisecond)
		n := <-result
		fmt.Printf("\twatcher() sleep: %v msec.\tdone %v\n", n, r)
		r++
		if r >= 4 {
			fmt.Printf("is done %v\n", r)
			for j := range quit {
				fmt.Printf("Init stop %v %#v \n", j, quit[j])
				if p.isDone[j] == false {
					fmt.Printf("\t WORKER Stop %v %#v\n", j, quit[j])
					quit[j] <- true
				} else {
					fmt.Printf("\t WORKER is not RUN %v\n", j)
				}
			}
			done <- true
			quitErrorHandler <- true
			return
		}
	}
}

func main() {

	var maxPool int = 5
	p := NewPool(maxPool)
	p.Run(work)

}
