package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FullName struct {
	firstname string
	lastname  string
}

func main() {
	//Create standard input
	reader := bufio.NewReader(os.Stdin)

	//file name
	fmt.Printf("Enter the name of the file: ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSuffix(fileName, "\n")
	fileName = strings.TrimSuffix(fileName, "\r")

	//read the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	names := make([]FullName, 0, 3)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		names = append(names, FullName{split[0], split[1]})

	}

	for _, n := range names {
		fmt.Println(n.firstname + " " + n.lastname)
	}
}
