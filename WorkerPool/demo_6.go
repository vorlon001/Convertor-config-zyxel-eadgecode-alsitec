package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
	
	"bufio"
	"bytes"
	"runtime/pprof"
	
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
	p := POOL{maxPool: maxPool ,
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

func (p *POOL) Run(isDone chan<- bool, work func(int, int) int) {

	for i := range p.Pool {
		fmt.Printf("run id %v %#v\n", i, p.Quit[i])
		go p.Worker(i, (*p).Quit[i], p.Error, p.Result, work)
	}
	go p.Watcher(p.Done, p.Quit, p.Result)
	<-p.Done
	isDone <-true
}

func (p *POOL) Worker(id int, quit chan bool, chError chan<- *workerError, result chan<- int, worker func(int, int) int) {
	defer panicRecoverWorker(chError, id, quit, &p.isDone)

	for {
		select {
		case <-quit:
			fmt.Printf("worker() id:%v is resived command QUIT!\n", id)
			return
		default:
			fmt.Printf("\tworker() id:%v\n", id )
			var n int = worker(id, 50)
			fmt.Printf("worker() id:%v\n", id )
			result <- n
			p.isDone[id] = true
			close(quit)
			return
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
				fmt.Println("STOP ERROR HANDLER WORKER")
				return
			case e := <-p.Error:
				fmt.Printf("ERROR in WORKER: \t%[1]v\n\t%#[1]v\n", e)
			}
		}
	}(quitErrorHandler)

	r := 0
	for {

		sum := func (array []bool) int {  
			result := 0  
			for _, v := range array {  
				if v==true {
					result++  
				}
			}  
			return result  
		}

		fmt.Printf("\t<<< %v %v \n", p.isDone, sum(p.isDone) )
		
		if sum(p.isDone)==p.maxPool {
			fmt.Printf("\t QUIT WAITER() \n")
			done <- true
			quitErrorHandler <- true
			return
		}
		
		n, ok  := <-result
		if ok==false {
			fmt.Printf("ERROR in Watcher \n")
		}
		fmt.Printf("\twatcher() sleep: %v msec.\tdone %v\n", n, r)
		r++
		if r >= 3 {
			fmt.Printf("is done %v %#v \n", r, p.isDone )
			go func (){
				for j := range quit {
					fmt.Printf(">>>>>>>>>>Init stop %v %#v \n", j, quit[j])
					if p.isDone[j] == false {
						fmt.Printf("\t WORKER Stop %v %#v\n", j, quit[j])					
						if _, ok := <-quit[j]; ok==true {
							quit[j] <- true
						}
					} else {
						fmt.Printf("\t WORKER is not RUN %v\n", j)
					}
				}
				done <- true
				quitErrorHandler <- true
			}()
			done <- true
			return
		}
	}
}

func main() {
	PrintMemUsage()
	
	isDone  := make(chan bool)
	var maxPool int = 5
	p := NewPool(maxPool)
	go p.Run( isDone , work)
	<-isDone 

	fmt.Printf("#############IS DONE")

	var b bytes.Buffer
	pprof.Lookup("goroutine").WriteTo(&b, 1)
	scanner := bufio.NewScanner(&b)
	for scanner.Scan() {
		t := scanner.Text()
		fmt.Printf(" %#v \n",t)
	}	
	
	time.Sleep(time.Duration(10000) * time.Millisecond)


	pprof.Lookup("goroutine").WriteTo(&b, 1)
	scanner = bufio.NewScanner(&b)
	for scanner.Scan() {
		t := scanner.Text()
		fmt.Printf(" %#v \n",t)
	}
	
	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()
}


// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}
