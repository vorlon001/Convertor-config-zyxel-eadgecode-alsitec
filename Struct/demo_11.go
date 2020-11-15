package main 
  
import "fmt"
import "reflect"

type Team struct {
	Ts string
	Sn string
}  

type team1 struct { 
	Team
	total int
} 



func (t1 *team1) Set(x interface{}) { 
	t1.total  = x.(int)
} 
  
type team2 struct { 
	Team
	name string
}



func (t1 *team2) Set(x interface{}) { 
	t1.name = x.(string)
}

func newObj( Set func (interface{}) )  func ( interface{} )   {
	return func (x interface{} ) {
		Set(x)
	}
}

func A(x interface{}, Set func (interface{}) ) {
	Set(x)
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

type H struct {
	f F
}
func (h H) getData() interface{} {
	return h.f.O
}

type H1 struct {
	H
}
func (h H1) getCall( f F) (ROCKET,func ( interface{} )) {
	newObj := func ( Set func (interface{}) )  func ( interface{} )   {
		return func (x interface{} ) {
			Set(x)
		}
	}
	h.f = f
	u := f.O.(*team1)
	return h,newObj(u.Set)
}

type H2 struct {
	H
}
func (h H2) getCall( f F) (ROCKET,func ( interface{} )) {
	newObj := func ( Set func (interface{}) )  func ( interface{} )   {
		return func (x interface{} ) {
			Set(x)
		}
	}
	h.f = f
	u := f.O.(*team2)
	return h,newObj(u.Set)
}

type ROCKET interface { 
	getCall( F) (ROCKET,func ( interface{} ))
	getData() interface{}
} 

func main() { 


	res1 := team1{total: 20} 
	f := newObj(res1.Set)
	f(1234)
	fmt.Printf(" %#v \n",res1) 
	f(1231234)
	fmt.Printf(" %#v \n",res1) 
	/*********************************/
	Z0 := func() F {
		return newObjs(team1{})		
	}
	Z1 := func(f F) (ROCKET,func ( interface{} )) {
		h0 := H1{}
		h1,f1 := h0.getCall(f)
		return h1,f1
		}
	Z2 := func(Z0 func() F, Z1 func(f F) (ROCKET,func ( interface{} ))) []interface{} {
		i0 :=  Z0()
		h1,f1 := Z1(i0)
		f1(1234)
		fmt.Printf(" %#v %#v %#v \n",h1.getData(),i0,i0.O)
		
		h2,f2 := Z1(Z0())
		f2(1234333)
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
		return newObjs(team2{})		
	}
	X1 := func(f F) (ROCKET,func ( interface{} )) {
		h0 := H2{}
		h1,f1 := h0.getCall(f)
		return h1,f1
		}
	X2 := func(Z0 func() F, Z1 func(f F) (ROCKET,func ( interface{} ))) []interface{} {
		i0 :=  Z0()
		h1,f1 := Z1(i0)
		f1("HEEEX XDFGWFBGHD")
		fmt.Printf(" %#v %#v %#v \n",h1.getData(),i0,i0.O)
		m := make([]interface{},0)
		m = append( m, h1.getData() )
		return m
	}
	
	mx := X2( X0, X1)
	mxs := mx[0].(*team2)
	fmt.Printf("2)>>>>>>>> %#v %#v  %#v %#v\n",mx[0],mx, mxs.Ts, mxs.name)
	/*********************************/
  
} 
