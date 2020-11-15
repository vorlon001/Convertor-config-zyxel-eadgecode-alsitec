package main 
  
import "fmt"

type employee interface { 
    Set(interface{})
} 
  

type team1 struct { 
    total int
} 

func (t1 *team1) Set(x int) { 
	t1.total  = x
} 
  
type team2 struct { 
    name string
} 

func (t1 *team2) Set(x string) { 
	t1.name = x
} 
  
func main() { 
  
	res1 := team1{total: 20} 
	res1.Set(234)
	fmt.Printf(" %#v \n",res1) 

	res2 := team2{ name: "2345234 def hgdf"} 
	res2.Set("sdfgf sdefgsdf gsdfgsdfg ")
	fmt.Printf(" %#v \n",res2) 
  
} 
