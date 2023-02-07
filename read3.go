package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type wholeName struct {
	fname string
	lname string
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var myPath string
	var listName []Name

	fmt.Print("Input path: ")
	input, _, _ := reader.ReadLine()
	myPath = string(input)
	fmt.Println("Your path: ", myPath)

	inputFile, err := os.Open(myPath)

	if err != nil {
		fmt.Println("Error: read file!")
	}

	reader2 := bufio.NewReader(inputFile)
	for {
		line, _, err := reader2.ReadLine()
		if err != nil {
			break
		}
		dataLine := strings.Split(string(line), " ")
		listName = append(listName, wholeName{dataLine[0], dataLine[1]})
	}

	fmt.Println("Output:")
	for i, x := range listName {
		fmt.Println("- Index:", i+1)
		fmt.Println("- First Name:", x.fname)
		fmt.Println("- Last Name:", x.lname)
		fmt.Println("-----------")
	}

}
