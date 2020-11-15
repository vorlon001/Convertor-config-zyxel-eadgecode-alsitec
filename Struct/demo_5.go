package main

import (
	"fmt"
)

type Access struct {
                Base       string   `yaml:"Base"`
                Group      string   `yaml:"Group"`
                PrivLvl    int      `yaml:"priv-lvl"`
                PERMIT     []string `yaml:"PERMIT,omitempty"`
                DENY       []string `yaml:"DENY,omitempty"`
        };
type User struct {
                Login    string   `yaml:"login"`
                Password string   `yaml:"password"`
                PrivLvl  int      `yaml:"priv-lvl"`
                IPAccess []string `yaml:"ip_access"`
                PERMIT   []string `yaml:"PERMIT"`
                DENY     []string `yaml:"DENY"`
        };

func get(param interface{}) {

    switch v := param.(type) { 
    default:
        fmt.Printf("unexpected type %T\n", v)
    case nil:
        fmt.Printf("interface{}(nil) type %T\n", v)
    case Access:
        fmt.Printf("main.Access type %T\n", v)
    case User:
        fmt.Printf("main.User type %T\n", v)
    } 
}

func yaml_cfg_get(param interface{}) interface{} {
	return 	param 
}
func main() {
	a := Access{}
	get(yaml_cfg_get(a));
	b := User{}	
	get(yaml_cfg_get(b));
	c := interface{}(nil)
	get(yaml_cfg_get(c));
}
