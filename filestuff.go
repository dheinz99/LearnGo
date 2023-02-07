package main

import (
	"fmt"
	"os"
	"strconv"
)

type Name1 struct {
	fName string
	lName string
}

func main() {
	var filename string
	var fileStore int
	fmt.Print("Enter the name of the file: ")
	fmt.Scan(&filename)
	fileStore, err := os.Open(filename)
	if err != nil {
		fmt.Println("error")
	}
	barr := make([]byte, 20)
	filestuff, err := filestore.Read(barr)
	filestore.Close()

	fmt.Printf("Json Object: %s", strconv.Itoa(filestuff))

}
