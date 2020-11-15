package main

import (
	"fmt"
)

func main() {

	t := make([]*string,0)
	t2 := make([]**string,0)
	t3 := make([]***string,0)
	t4 := make([]****string,0)
	s1  := "atwfasdfsdafs"
	s2  := &s1
	s3  := &s2
	s4  := &s3
	t = append(t,&s1);
	t2 = append(t2,&s2);
	t3 = append(t3,&s3);
	t4 = append(t4,&s4);
	fmt.Printf("%v %v \n", t , *t[0])
	r2 := *t2[0]
	
	r31 := *t3[0]
	r32 := *r31
	r33 := *r32

	r41 := *t4[0]
	r42 := *r41
	r43 := *r42
	r44 := *r43
	fmt.Printf("%v %v  %v %v %v %v %v\n", t2 , *t2[0], *r2,  t3, r33, t4,r44)
	
	t8 := &t
	t9 := &t8
	fmt.Printf("%#v %#v \n", t8, t9 )
	s8 := *t8
	fmt.Printf("%#v %#v  %#v\n", *t8, *s8[0], *t9 )	

	s91 := *t9
	s92 := *s91
	//s92 := *s91
	fmt.Printf("%#v  %#v %#v\n", *s91, *s92[0], *t9 )
}
