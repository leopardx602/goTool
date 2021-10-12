package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName01"` // must uppercase ,only for json output
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var structData Serverslice
	jsonStr := `{"servers":[{"serverName":"Chen_VPN","serverIP":"127.0.0.1"},{"serverName":"Ting_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(jsonStr), &structData) // json string -> struct
	fmt.Println(structData)
	fmt.Println(structData.Servers[0].ServerIP)

	jsonData, _ := json.Marshal(structData) // struct -> json
	jsonStr = string(jsonData)              // json -> json string
	fmt.Println(jsonStr)
}
