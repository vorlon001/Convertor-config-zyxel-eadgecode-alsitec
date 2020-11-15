package main

import (
	"fmt"
	"time"
	"reflect"
	"sync"
)


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

func (h ROCKET_OBJ) getCall( f ROCKET_PROXY , K func (h *interface{}) func ( string, interface{} ) ) (ROCKET,func ( string, interface{} )) {
	newObj := func ( Set func (string,interface{}) )  func ( string, interface{} )   {
		return func (id string, x interface{} ) {
			Set(id,x)
		}
	}
	h.f = f
	k := K(&f.O)
	return h, newObj(k)
}



type ROCKET interface { 
	getCall( ROCKET_PROXY, func (h *interface{}) func ( string, interface{} )) (ROCKET,func ( string, interface{} ))
	getData() interface{}
} 

type WORKER struct {
	wg *sync.WaitGroup
	ping chan *interface{}
	stop chan bool
	doneping chan bool
	donestop chan bool
}
func NewWORKER() *WORKER{
	w := WORKER{}
	w.wg = &sync.WaitGroup{}
	w.ping = make(chan *interface{},30)
	w.stop = make(chan bool)
	w.doneping = make(chan bool)
	w.donestop = make(chan bool)
	return &w
}

func (w WORKER) pinger(Z0 func() ROCKET_PROXY , Z1 func(f ROCKET_PROXY ) (ROCKET,func ( string, interface{} ))) {
	defer w.wg.Done()
	j :=0
	for {
		if j>30 {
			fmt.Println("pinger - init send stop and close")
			for { 
				if len(w.ping)==0 {
					break
				}
				time.Sleep(1000 * time.Millisecond)
			}
			w.stop <- true
			close(w.doneping)
			fmt.Println("pinger - stop is done")
			return
		}

		time.Sleep(100 * time.Millisecond)

		i0 :=  Z0()
		h1,f1 := Z1(i0)
		f1("Total",10000+j)
		f1("Sn", fmt.Sprintf("SNSNSN%d",10000+j) )
		fmt.Printf("ping %d %#v \n",j,h1.getData())

		var s interface{}
		s = h1.getData()	
		w.ping<- &s
		j++		
	}
}

func (w WORKER) ponger(Z2 func(*interface{}) []interface{} ) {
	defer w.wg.Done()
	for {
		select {
			case val := <-w.ping:
				l := Z2(val)
				fmt.Printf("pong %#v %#v \n", *val , l )
				time.Sleep(300 * time.Millisecond)
			case <-w.stop:
				fmt.Println("ponger - stop is done")
				close(w.donestop)
				return
		}
	}
}

func (w WORKER) Run(Z *ROCKET_Model ) {
	w.wg.Add(1)
	go w.pinger( (*Z).NewNode, (*Z).InitSetFunc)
	w.wg.Add(1)
	go w.ponger( (*Z).RunLogic)
	w.wg.Wait()
}

func (w WORKER) Wait() {
	w.wg.Wait()
	close(w.ping )
	close(w.stop )
	_ = [2]bool{<-w.doneping , <-w.donestop}
}

type ROCKET_Model interface {
	NewNode() ROCKET_PROXY
	InitSetFunc(f ROCKET_PROXY ) (ROCKET,func ( string, interface{} ))
	RunLogic(i *interface{}) []interface{}
}
type ROCKET_NODE struct {
}

func (rn ROCKET_NODE) newObjs(obj interface{}) ROCKET_PROXY  {
	to_struct_ptr := func (obj interface{}) interface{} {
		vp := reflect.New(reflect.TypeOf(obj))
		vp.Elem().Set(reflect.ValueOf(obj))
		return vp.Interface()
	}
	f := to_struct_ptr( obj )
	return ROCKET_PROXY{ O:f }
}

/****************************************************/

type Node_Block struct { 
	Node
	total int
} 

func (t1 *Node_Block) Set(id string, x interface{}) { 
	switch id {
		case "Total":
			t1.total  = x.(int)
		case "Sn":
			t1.Sn = x.(string)
	}
} 

type Node_Model struct {
	ROCKET_NODE
}

func (nm Node_Model) NewNode() ROCKET_PROXY  {
	return nm.newObjs(Node_Block{})		
}
func (nm Node_Model) InitSetFunc(f ROCKET_PROXY ) (ROCKET,func ( string, interface{} )) {
	h0 := ROCKET_OBJ{}
	K := func (h *interface{}) func ( string, interface{} ) {
		u := f.O.(*Node_Block)
		return u.Set
	}
	h1,f1 := h0.getCall(f, K)
	return h1,f1
}
func (nm Node_Model) RunLogic(i *interface{}) []interface{} {
	h := (*i).(*Node_Block);
	r := make([]interface{},0)
	r = append( r, h.total)
	r = append( r, h.Sn)
	return r
}
func main() {

	var U ROCKET_Model 
	U = Node_Model{}
	w := NewWORKER()
	w.Run(&U)
	w.Wait()
}

