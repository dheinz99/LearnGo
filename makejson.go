package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var addr string

	fmt.Print("Please Enter a name: ")
	fmt.Scan(&name)

	fmt.Print("Please enter an address: ")
	fmt.Scan(&addr)

	nad := map[string]string{
		"name":    name,
		"address": addr,
	}
	barr, _ := json.Marshal(&nad)

	fmt.Printf("Json Object: %s", barr)
}
