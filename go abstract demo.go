package main 
  
import "fmt"
import "reflect"

type Team struct {
	Ts string
	Sn string
}  

type Team1 struct { 
	Team
	total int
} 



func (t1 *Team1) Set(id string, x interface{}) { 
	switch id {
		case "Total":
			t1.total  = x.(int)
	}
} 
  
type Team2 struct { 
	Team
	name string
}



func (t1 *Team2) Set(id string, x interface{}) {
	switch id {
		case "Name":
			t1.name = x.(string)
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

type H1 struct {
	ROCKET_OBJ 
}
func (h H1) getCall( f F) (ROCKET,func ( string, interface{} )) {
	newObj := func ( Set func (string,interface{}) )  func ( string, interface{} )   {
		return func (id string, x interface{} ) {
			Set(id,x)
		}
	}
	h.f = f
	u := f.O.(*Team1)
	return h,newObj(u.Set)
}

type H2 struct {
	ROCKET_OBJ 
}
func (h H2) getCall( f F) (ROCKET,func ( string, interface{} )) {
	newObj := func ( Set func (string,interface{}) )  func ( string, interface{} )   {
		return func (id string, x interface{} ) {
			Set(id,x)
		}
	}
	h.f = f
	u := f.O.(*Team2)
	return h,newObj(u.Set)
}

type ROCKET interface { 
	getCall( F) (ROCKET,func ( string, interface{} ))
	getData() interface{}
} 

func main() { 

	/*********************************/
	Z0 := func() F {
		return newObjs(Team1{})		
	}
	Z1 := func(f F) (ROCKET,func ( string, interface{} )) {
		h0 := H1{}
		h1,f1 := h0.getCall(f)
		return h1,f1
		}
	Z2 := func(Z0 func() F, Z1 func(f F) (ROCKET,func ( string, interface{} ))) []interface{} {
		i0 :=  Z0()
		h1,f1 := Z1(i0)
		f1("Total",1234)
		fmt.Printf(" %#v %#v %#v \n",h1.getData(),i0,i0.O)
		
		h2,f2 := Z1(Z0())
		f2("Total",1234333)
		fmt.Printf(" %#v \n",h2.getData())

		
		m := make([]interface{},0)
		m = append( m, h1.getData() )
		m = append( m, h2.getData() )
		return m
	}
	
	mz := Z2( Z0, Z1)
	fmt.Printf("1)>>>>>>>> %#v %#v \n",mz[0],mz[1])
	/*********************************/

	X0 := func() F {
		return newObjs(Team2{})		
	}
	X1 := func(f F) (ROCKET,func ( string, interface{} )) {
		h0 := H2{}
		h1,f1 := h0.getCall(f)
		return h1,f1
		}
	X2 := func(Z0 func() F, Z1 func(f F) (ROCKET,func ( string, interface{} ))) []interface{} {
		h1,f1 := Z1(Z0())
		f1("Name","HEEEX XDFGWFBGHD")
		fmt.Printf(" %#v \n",h1.getData())
		m := make([]interface{},0)
		m = append( m, h1.getData() )
		return m
	}
	
	mx := X2( X0, X1)
	mxs := mx[0].(*Team2)
	fmt.Printf("2)>>>>>>>> %#v %#v  %#v %#v\n",mx[0],mx, mxs.Ts, mxs.name)
	/*********************************/
  
} 
