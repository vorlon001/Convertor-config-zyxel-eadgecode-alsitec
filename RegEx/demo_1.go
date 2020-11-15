package main

import (
    "regexp"
    "fmt"
)
func Verify_Cmd(permit []string, cmd string) bool {
	var Find_Cmd =func(rgx string, cmd string) bool{
		var re = regexp.MustCompile(rgx)
		if a := re.FindAllString(cmd, -1);len(a)==1 {
			return true
	    	}
		return false
	}
	var status = false
	for _,v:= range permit  {
    		if Find_Cmd(v,cmd)==true {
			status = true
		}
	}
	return status 
}
func main() {
    var e = []string{"telnet","configure","terminal","show","write"}
//    var e = []string{"telnet","configure","terminal","show"}
    var str = "  write <cr>"
    fmt.Println("found at index", Verify_Cmd(e,str))
}
