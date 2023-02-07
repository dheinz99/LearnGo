package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Print("Please Enter String Value: ")
	fmt.Scan(&str)

	switch {
	case strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a"):
		fmt.Println("Found!")
	default:
		fmt.Println("Not Found!")
	}
}
