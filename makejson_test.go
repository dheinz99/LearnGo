package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	in1 := bufio.NewReader(os.Stdin)
	d1 := make(map[string]string)

	fmt.Println("Enter the name :")
	tempvar1, _, err_x := in1.ReadLine()
	if err_x != nil {
		fmt.Println(err_x)
	}
	d1["name"] = string(tempvar1)
	fmt.Println("Enter the address :")
	tempvar1, _, err_y := in1.ReadLine()
	if err_y != nil {
		fmt.Println(err_y)
	}
	d1["address"] = string(tempvar1)
	jsn_dl, err_jsn := json.Marshal(d1)
	if err_jsn != nil {
		fmt.Println(err_jsn)
	}
	//    fmt.Println(jsn_dl)
	fmt.Println(string(jsn_dl))
}
