package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
	//"errors"
	"log"
)
func panicRecover() {
        if r := recover(); r != nil {
                log.Printf("Internal error: %v", r)
                buf := make([]byte, 1<<16)
                stackSize := runtime.Stack(buf, true)
		log.Printf("--------------------------------------------------------------------------------")
                log.Printf("Internal error: %s\n", string(buf[0:stackSize]))
		log.Printf("--------------------------------------------------------------------------------")
        }
    }
/*********************************************************************************************************************/
type HandlerError struct {
	Err	string
	Dump	string
	WorkerError *workerError
}
func (e *HandlerError) Error() string {
	return fmt.Sprintf("Err:%[1]s", e.Err)
}
/*********************************************************************************************************************/
type workerError struct {
	pingChanId int
	stopChanId int
	doneChanId int
	quitChanId int
	Err	string
	Dump	string
}
func (e *workerError) Error() string {
	return fmt.Sprintf("workerError:%[1]d stopChanId:%[2]d doneChanId:%[3]d quitChanId:%[4]d Err:%[5]s",e.pingChanId,e.stopChanId,e.doneChanId,e.quitChanId, e.Err)
}
/*********************************************************************************************************************/

func panicRecoverWorker(i chan *workerError, pingChanId int, stopChanId int, doneChanId int,quitChanId int ) {
        if r := recover(); r != nil {
                log.Printf("Run panicRecoverWorker(), Internal error: %v", r)
                buf := make([]byte, 1<<16)
                stackSize := runtime.Stack(buf, true)
		i<-&workerError{
				pingChanId:pingChanId,
				stopChanId: stopChanId,
				doneChanId: doneChanId,
				quitChanId: quitChanId,
				Dump: string(buf[0:stackSize]),
				Err: fmt.Sprintf("Internal error: %v", r),
		}
	 }
 }

	
type Node struct {
	Ts string
	Sn string
}

type ROCKET_PROXY struct {
	O interface{}
}

type ROCKET_OBJ struct {
	f ROCKET_PROXY
}

func (h ROCKET_OBJ) getData() interface{} {
	return h.f.O
}

func (h ROCKET_OBJ) getCall(f ROCKET_PROXY, K func(h *interface{}) func(string, interface{})) (ROCKET, func(string, interface{})) {
	newObj := func(Set func(string, interface{})) func(string, interface{}) {
		return func(id string, x interface{}) {
			Set(id, x)
		}
	}
	h.f = f
	k := K(&f.O)
	return h, newObj(k)
}

type ROCKET interface {
	getCall(ROCKET_PROXY, func(h *interface{}) func(string, interface{})) (ROCKET, func(string, interface{}))
	getData() interface{}
}

type WORKER struct {
	Debug   bool
	Workers int
	Z       *ROCKET_Model
	ping    *[]chan *interface{}
	stop    *[]chan bool
	done    *[]chan *workerError
	quit	*[]chan bool
}

func NewWORKER(max int, Lengh int, Debug bool) *WORKER {
	w := WORKER{}
	w.Debug = Debug
	initChanBool := func(max int) *[]chan bool {
		c := make([]chan bool, max)
		for k, _ := range c {
			c[k] = make(chan bool)
		}
		return &c
	}
	initChanPing := func(max int) *[]chan *interface{} {
		c := make([]chan *interface{}, max)
		for k, _ := range c {
			c[k] = make(chan *interface{}, Lengh)
		}
		return &c
	}

	initChanWorkerError := func(max int) *[]chan *workerError {
		c := make([]chan *workerError, max)
		for k, _ := range c {
			c[k] = make(chan *workerError)
		}
		return &c
	}	
	
	w.Workers = max
	w.ping = initChanPing(w.Workers)
	w.stop = initChanBool(w.Workers)
	w.done = initChanWorkerError(w.Workers)
	w.quit = initChanBool(w.Workers)
	return &w
}

func (w WORKER) pinger(pingChanId int, stopPongChanId int, doneChanId int, quitChanId int) {
	defer panicRecoverWorker((*w.done)[doneChanId], pingChanId, stopPongChanId, doneChanId,quitChanId	)
	j := 0
	
	for {
	        select {
		        case <- (*w.quit)[quitChanId]:
				(*w.done)[doneChanId] <- nil
				return
		default:
			if j > 3 {
				//panic("33333333333")
			}
			if j > 30 {
				if w.Debug == true {
					log.Println("pinger - init send stop and close")
				}
				for {
					if len((*w.ping)[pingChanId]) == 0 {
						break
					}
					time.Sleep(1000 * time.Millisecond)
				}
				(*w.stop)[stopPongChanId] <- true
				(*w.done)[doneChanId] <- nil
				if w.Debug == true {
					log.Println("pinger - stop is done", doneChanId)
				}
				return
			}
			time.Sleep(100 * time.Millisecond)
			run := func() *interface{}{
				i0 := (*w.Z).NewNode()
				h1, f1 := (*w.Z).InitSetFunc(i0)
				f1("Total", 10000+j)
				f1("Sn", fmt.Sprintf("SNSNSN%d", 10000+j))
				log.Printf("ping %d %#v \n", j, h1.getData())
				var s interface{}
				s = h1.getData()
				return &s
			}
			(*w.ping)[pingChanId] <- run()
			j++
		}
	}
}

func (w WORKER) ponger(pingChanId int, stopMyChanId int, doneChanId int,quitChanId int) {
	defer panicRecoverWorker((*w.done)[doneChanId], pingChanId, stopMyChanId, doneChanId,quitChanId	)
	for {
		select {
		case val := <-(*w.ping)[pingChanId]:
			run := func() *[]interface{}{
				l := (*w.Z).RunLogic(val)
				return &l
			}
			l := run();
			if w.Debug == true {
				log.Printf("pong %#v %#v \n", *val, *l)
			}
			time.Sleep(300 * time.Millisecond)
		case <-(*w.stop)[stopMyChanId]:
			if w.Debug == true {
				log.Println("ponger - stop is done", doneChanId)
			}
			(*w.done)[doneChanId] <- nil
			return
		case <- (*w.quit)[quitChanId]:
			(*w.done)[doneChanId] <- nil
			return
		}
	}
}

func (w WORKER) Run(Z *ROCKET_Model) {
	w.Z = Z
	go w.pinger(0, 0, 0, 0)
	go w.ponger(0, 0, 1, 1)

	go w.pinger(1, 1, 2, 2)
	go w.ponger(1, 1, 3, 3)
	
}

func (w WORKER) Wait() ( error *HandlerError){
	defer func () {
		        if r := recover(); r != nil {
	        	        log.Printf("Internal error: %v", r)
                		buf := make([]byte, 1<<16)
        		        stackSize := runtime.Stack(buf, true)
				error = &HandlerError{
					Dump: string(buf[0:stackSize]),
					Err: fmt.Sprintf("Error in Wait. Internal error: %v", r),
					WorkerError: nil,
				}
        		}
		}()
	Join := func(inputs ...chan *workerError) *workerError{
		var done  chan *workerError
		done = make(chan *workerError)
		doneWorker := func(done chan *workerError, i chan *workerError, k int) {
			if w.Debug == true {
				log.Println("Wait Join(): Wait CLOSE WORKER", i, k)
			}
			switch m := <-i; m {
				case nil:
					if w.Debug == true {
						log.Println("Wait Join():STOP WORKER", k, " DONE CORRECT")
					}
				default:					
					if w.Debug == true {
						log.Printf("Wait Join():STOP WORKER %v %#v DONE NOT CORRECT\n", k, m )
					}
					for l, q := range *w.quit {
						if l!=k {
							q<-true
							log.Println("ERROR STOP WORKER", l, " DONE NOT CORRECT")
						}
					}
					done<-m
					close(i)
					return 
			}
			done<-nil
			close(i)
		}
		for k, num := range inputs {			
			go doneWorker(done, num, k)

		}
		doneWorkers := 0
		for {
			s:=<-done;
			if w.Debug == true {
				log.Printf("WORKER IS DONE STATUS OK\n")
			}
			doneWorkers++
			if s!=nil {
				if w.Debug == true {
					log.Printf("WAIT JOIN WORKER HANDLER ERROR: %#v \n",s);
				}
				close(done)
				return s
				
			}
			if doneWorkers==w.Workers {
				break
			}
		}
		close(done)
		return nil
	}
	ok := Join((*w.done)...)
	error = nil
	if ok!=nil {
		error = &HandlerError{
			Dump: ok.Dump,
			Err: fmt.Sprintf("ERROR Worker HANDLER: %s", ok.Err),
			WorkerError: ok,
      		}
	}
	for _, v := range *w.ping {
		close(v)
	}
	for _, v := range *w.stop {
		close(v)
	}
	return error
}

type ROCKET_Model interface {
	NewNode() ROCKET_PROXY
	InitSetFunc(f ROCKET_PROXY) (ROCKET, func(string, interface{}))
	RunLogic(i *interface{}) []interface{}
}
type ROCKET_NODE struct {
}

func (rn ROCKET_NODE) newObjs(obj interface{}) ROCKET_PROXY {
	to_struct_ptr := func(obj interface{}) interface{} {
		vp := reflect.New(reflect.TypeOf(obj))
		vp.Elem().Set(reflect.ValueOf(obj))
		return vp.Interface()
	}
	f := to_struct_ptr(obj)
	return ROCKET_PROXY{O: f}
}

/****************************************************/

type Node_Block struct {
	Node
	total int
}

func (t1 *Node_Block) Set(id string, x interface{}) {
	switch id {
	case "Total":
		t1.total = x.(int)
	case "Sn":
		t1.Sn = x.(string)
	}
}

type Node_Model struct {
	ROCKET_NODE
}

func (nm Node_Model) NewNode() ROCKET_PROXY {
	return nm.newObjs(Node_Block{})
}
func (nm Node_Model) InitSetFunc(f ROCKET_PROXY) (ROCKET, func(string, interface{})) {
	h0 := ROCKET_OBJ{}
	K := func(h *interface{}) func(string, interface{}) {
		u := f.O.(*Node_Block)
		return u.Set
	}
	h1, f1 := h0.getCall(f, K)
	return h1, f1
}
func (nm Node_Model) RunLogic(i *interface{}) []interface{} {
	h := (*i).(*Node_Block)
	r := make([]interface{}, 0)
	r = append(r, h.total)
	r = append(r, h.Sn)
	return r
}
/****************************************************/
func main() {
	defer panicRecover() 
	var U ROCKET_Model
	U = Node_Model{}

	t := reflect.TypeOf(U)
	log.Printf("STRUCT: %#v \n", t.Name())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		log.Printf("\t Method %#v %#v \n", m.Name, m.Func)
	}
	log.Printf("------------------------------\n")
	w := NewWORKER(4, 30, true)
	w.Run(&U)
	
	ok := w.Wait();
	if ok!=nil {
	 	log.Printf(" %#v %#v \n", ok.Err, ok.Dump)
		if ok.WorkerError!=nil {
			fmt.Printf("WorkerError:%s \n DUMP:%s \n",ok.WorkerError.Error(), ok.WorkerError.Dump)
		}
	}
}
