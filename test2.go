package main

import "fmt"

func main() {
	var myNumber float64

	fmt.Print("Please enter a floating point number: ")
	fmt.Scan(&myNumber)
	finalNum := int(myNumber)
	fmt.Println("Truncated: ", finalNum)

}
