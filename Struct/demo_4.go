package main
import "fmt"

func Validate(a *c,h int ) {
	fmt.Println(*a)
	c := *a
	c[1]=h
}
func Validate2(a *[] int,h int ) {
	fmt.Println(*a)
	c := *a
	c[1]=h
}
type H struct{  Config *c }
type H2 struct{  Config *[] int }
type c [] int;

var d = &c{4,4,4,4}

func main() {
	fmt.Printf("%T %#v\n",d,d);
	Validate(d,2345);
	fmt.Printf("%T %#v\n",d,d);
	z := &H{Config: d}
	fmt.Printf("%T %#v\n",z.Config,z.Config);
	Validate(d,356);
	fmt.Printf("%T %#v\n",z.Config,z.Config);
	m := &([] int{4,4,4,4})
	z2 := &H2{Config: m}
	z3 := &H2{Config: m}
	fmt.Printf("%T %#v\n",z2.Config,z2.Config);
	fmt.Printf("%T %#v\n",z3.Config,z3.Config);
	Validate2(m,32356);
	fmt.Printf("%T %#v\n",z2.Config,z2.Config);
	fmt.Printf("%T %#v\n",z3.Config,z3.Config);
	
	f := *z3.Config;
	f[2]=2342534;
	fmt.Printf(">%T %#v\n",z2.Config,z2.Config);
	fmt.Printf(">%T %#v\n",z3.Config,z3.Config);
	fmt.Printf(">%T %#v\n",m,m);
	

	e := 2342;
	v := &e
	fmt.Printf("%T %#v\n",e,e);
	fmt.Printf("%T %#v\n",v,v);
	fmt.Printf("%T %#v\n",*v,*v);

}
