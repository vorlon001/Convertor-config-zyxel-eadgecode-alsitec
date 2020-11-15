package main

import (
    	"regexp"
    	"fmt"
	"time"
	"strconv"
)

func Paring_Args_Request(Args []string) map[string]string {

	var parsing  = func (re *regexp.Regexp, Args []string) map[string]string {
		s := make(map[string]string,len(Args ))
		for _,v := range Args  {
			if a := re.FindAllString(v, -1);len(a)==2 {
				if  a[0]=="start_time" {
					if s, err := strconv.ParseInt(a[1], 10, 64); err == nil {
						a[1] = fmt.Sprintf("%s\n", time.Unix(s, 0))
					}				
				}
				s[string(a[0])]=string(a[1])
			}
		}
		return s
	}


	var re = regexp.MustCompile(`[\w\-\<\>\s]+`)
	return parsing (re, Args)
	}
	
func main() {
	
	Args := []string{"task_id=406", "start_time=1593850973", "timezone=YEKT", "service=shell", "priv-lvl=15", "cmd=write <cr>"}
	arg := Paring_Args_Request( Args)
	fmt.Printf("%T %v\n",arg,arg)
	for k,v := range arg {
		fmt.Println(k,v)
	}
}
