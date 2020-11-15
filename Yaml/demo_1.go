package main

import (
        "fmt"
        "log"

        "gopkg.in/yaml.v2"
)

var data = `
PID: "PATH"
BIND: 0.0.0.0
PORT: 49
LOG:
  DEBUG:
      MODE: TRUE
  SYSLOG:
      ENABLE: TRUE
      UP: 1.1.1.1
      PORT: 49
  FILE:
      ENABLE: TRUE
      NAME: "FILE.LOG"
LDAP:
  Base:   "BASE"       
  Host:   "HOST"
  Port:   389
  UseSSL: 0
  BindDN: "BindDN"
  BindPassword: "BindPassword"
  UserFilter: "UserFilter"
  GroupFilter: "GroupFilter"
  Attributes:  
    - "Attributes1" #LDAP GROUP LEVEL 1
    - "Attributes2" #LDAP GROUP LEVEL 15 limitted
    - "Attributes3" #LDAP GROUP LEVEL 15
ACCESS: # ACCESS for LDAP USERS
    - Attributes: Attributes1
      priv-lvl: 1
    - Attributes: Attributes2
      priv-lvl: 15
      PERMIT:
        - "telnet"
        - "configure terminal"
        - "show"
      DENY:
        - "reboot"
        - "dir"
    - Attributes: Attributes3
      priv-lvl: 15
banner:
  login_banner: "LOGIN:"
  password_banner: "PASSWORD:"
  banner: "sdfgsDFgsdfgsdfgsdf \nsdfG"
  banner_accept: "sdfgsDFgsdfgsdfgsdf \nsdfG"
  banner_reject: "sdfgsDFgsdfgsdfgsdf \nsdfG"
access:
  - network: 10.0.0.0/24
    token: fsdasdf3
  - network: 10.0.1.0/24
    token: fsdasdfw
user: # if need
  - login: vorlon3
    password: sdfasdfasd_need_argon_password_and_util_pwdgen
    priv-lvl: 15
    ip_access: 
      - 10.0.0.1/32 
      - 10.0.0.10/32
      - 10.0.1.0/24
    PERMIT:
      - "telnet"
      - "configure terminal"
      - "show"
    DENY:
      - "reboot"
      - "dir"
`  



type AutoGenerated struct {
	PID  string `yaml:"PID"`
	BIND string `yaml:"BIND"`
	PORT int    `yaml:"PORT"`
	LOG  struct {
		DEBUG struct {
			MODE bool `yaml:"MODE"`
		} `yaml:"DEBUG"`
		SYSLOG struct {
			ENABLE bool   `yaml:"ENABLE"`
			UP     string `yaml:"UP"`
			PORT   int    `yaml:"PORT"`
		} `yaml:"SYSLOG"`
		FILE struct {
			ENABLE bool   `yaml:"ENABLE"`
			NAME   string `yaml:"NAME"`
		} `yaml:"FILE"`
	} `yaml:"LOG"`
	LDAP struct {
		Base         string   `yaml:"Base"`
		Host         string   `yaml:"Host"`
		Port         int      `yaml:"Port"`
		UseSSL       int      `yaml:"UseSSL"`
		BindDN       string   `yaml:"BindDN"`
		BindPassword string   `yaml:"BindPassword"`
		UserFilter   string   `yaml:"UserFilter"`
		GroupFilter  string   `yaml:"GroupFilter"`
		Attributes   []string `yaml:"Attributes"`
	} `yaml:"LDAP"`
	ACCESS []struct {
		Attributes string   `yaml:"Attributes"`
		PrivLvl    int      `yaml:"priv-lvl"`
		PERMIT     []string `yaml:"PERMIT,omitempty"`
		DENY       []string `yaml:"DENY,omitempty"`
	} `yaml:"ACCESS"`
	Banner struct {
		LoginBanner    string `yaml:"login_banner"`
		PasswordBanner string `yaml:"password_banner"`
		Banner         string `yaml:"banner"`
		BannerAccept   string `yaml:"banner_accept"`
		BannerReject   string `yaml:"banner_reject"`
	} `yaml:"banner"`
	Access []struct {
		Network string `yaml:"network"`
		Token   string `yaml:"token"`
	} `yaml:"access"`
	User []struct {
		Login    string   `yaml:"login"`
		Password string   `yaml:"password"`
		PrivLvl  int      `yaml:"priv-lvl"`
		IPAccess []string `yaml:"ip_access"`
		PERMIT   []string `yaml:"PERMIT"`
		DENY     []string `yaml:"DENY"`
	} `yaml:"user"`
}

func Var_dump(expression ...interface{} ) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

func main() {
        t := AutoGenerated{}
        fmt.Println(fmt.Sprintf("%v:%v","0.0.0.0",44))
        err := yaml.Unmarshal([]byte(data), &t)
        if err != nil {
                log.Fatalf("error: %v", err)
        }
        fmt.Printf("--- t:\n%T %v\n\n", t,t)
        fmt.Printf("--- t:\n%T %v\n\n", t.LDAP.Base,t.LDAP.Base)
	Var_dump(t)
 
}
