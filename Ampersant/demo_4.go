package main

import (
	"fmt"
)

func intSet1(a []int) {
	a[1]++
	a[1]++
}
func intSet2(a *[]int) {
	(*a)[1]++
	(*a)[1]++
}
func main() {
	slice1:= []string{"a1","a2","a3","a4","a5","a6","a7","a8","a9","a10"}
	slice2:=  &slice1
	var slice3 []string
	slice3= (*slice2)[4:6]
	slice3[0]="2q34123"
	
	var slice4 []string
	slice4= slice1[4:6]
	slice4[1]="2q3412333sfd"
	
	fmt.Printf(" %#v  \n% #v  \n %#v \n %#v \n",slice1,slice2,slice3,slice4)
	
	slice10:= []int{1,2,3,4,4,5,6,7,8,9,10}

	slice20:=  &slice10
	var slice30 []int
	slice30= (*slice20)[4:6]
	slice30[0]++
	intSet1(slice30);
	intSet2(&slice10);	
	
	fmt.Printf(" %#v  \n% #v  \n %#v \n",slice10,slice20,slice30)

}
