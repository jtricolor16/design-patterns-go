package main

import "fmt"

func main() {
	request := "fluminense:campeão,flamengo:rebaixado,vasco:nada,botafogo:nada"
	c := NewConvertStringToJson()
	fmt.Println(c.KeyValueToJson(request))
}
