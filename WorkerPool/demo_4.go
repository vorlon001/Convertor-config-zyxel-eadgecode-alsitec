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
  
type F struct {
	O interface{} 
}
func newObjs(obj interface{}) F {
	to_struct_ptr := func (obj interface{}) interface{} {
		vp := reflect.New(reflect.TypeOf(obj))
		vp.Elem().Set(reflect.ValueOf(obj))
		return vp.Interface()
	}
	f := to_struct_ptr( obj )
	return F{ O:f }
}

type ROCKET_OBJ struct {
	f F
}
func (h ROCKET_OBJ) getData() interface{} {
	return h.f.O
}

func (h ROCKET_OBJ) getCall( f F , K func (h *interface{}) func ( string, interface{} ) ) (ROCKET,func ( string, interface{} )) {
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
	getCall( F, func (h *interface{}) func ( string, interface{} )) (ROCKET,func ( string, interface{} ))
	getData() interface{}
} 


func pinger(pinger chan<- *interface{}, stop chan<- bool, done chan bool, wg *sync.WaitGroup,Z0 func() F, Z1 func(f F) (ROCKET,func ( string, interface{} ))) {
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

		time.Sleep(100 * time.Millisecond)

		i0 :=  Z0()
		h1,f1 := Z1(i0)
		f1("Total",10000+j)
		f1("Sn", fmt.Sprintf("SNSNSN%d",10000+j) )
		fmt.Printf("ping %d %#v \n",j,h1.getData())

		var s interface{}
		s = h1.getData()	
		pinger <- &s
		j++		
	}
}


func ponger(pinger <-chan *interface{}, stop <-chan bool, done chan bool, wg *sync.WaitGroup, Z2 func(*interface{}) []interface{} ) {
	defer wg.Done()
	for {
		select {
			case val := <-pinger:
				l := Z2(val)
				fmt.Printf("pong %#v %#v \n", *val , l )
				time.Sleep(300 * time.Millisecond)
			case <-stop:
				fmt.Println("ponger - stop is done")
				close(done)
				return
		}
	}
}

func main() {


	Z0 := func() F {
		return newObjs(Node_Block{})		
	}
	Z1 := func(f F) (ROCKET,func ( string, interface{} )) {
		h0 := ROCKET_OBJ{}
		K := func (h *interface{}) func ( string, interface{} ) {
			u := f.O.(*Node_Block)
			return u.Set
		}
		h1,f1 := h0.getCall(f, K)
		return h1,f1
		}
	Z2 := func(i *interface{}) []interface{} {
		h := (*i).(*Node_Block);
		r := make([]interface{},0)
		r = append( r, h.total)
		r = append( r, h.Sn)
		return r
	}
	
	var wg sync.WaitGroup
	ping := make(chan *interface{},30)
	stop := make(chan bool)
	doneping := make(chan bool)
	donestop := make(chan bool)
	
	wg.Add(1)
	go pinger(ping, stop, doneping, &wg, Z0, Z1)
	wg.Add(1)
	go ponger(ping, stop, donestop, &wg,Z2)
	wg.Wait()
	close(ping )
	close(stop )
	_ = [2]bool{<-doneping , <-donestop}
}

