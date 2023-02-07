package main

import "fmt"

func main() {
	var f float64

	fmt.Print("Enter a float value: ")
	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		var i int = int(f)
		fmt.Println("Float is truncated to: ", i)
	}
}
