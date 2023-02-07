package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Welcome, please introduce a character string and I'll tell you if I found it to comply with the requisites")
	fmt.Scanln(&input)
	i := strings.Index(strings.ToLower(input), "i")
	a := strings.Index(strings.ToLower(input), "a")
	n := strings.Index(strings.ToLower(input), "n")
	if (i == 0) && (n == len(input)-1) && (a != -1) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
