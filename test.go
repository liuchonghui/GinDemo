package main

import (
	"encoding/json"
	"fmt"
)

type Pl struct {
	Id string `json:"id"`
	Md5   string `json:"md5"`
}

type Serverslice struct {
	Servers []Pl `json:"servers"`
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Pl{Id: "Shanghai_VPN", Md5: "127.0.0.1"})
	s.Servers = append(s.Servers, Pl{Id: "Beijing_VPN", Md5: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}