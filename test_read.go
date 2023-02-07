package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	maxLength = 20
)

type Name struct {
	fname string
	lname string
}

func (n *Name) Set(fname string, lname string) {
	n.fname = fname
	n.lname = lname

	if len(fname) > maxLength {
		n.fname = fname[:maxLength]
	}

	if len(lname) > maxLength {
		n.lname = lname[:maxLength]
	}
}

func main() {
	var filename string

	var names []Name

	fmt.Println("Please, enter the filename: ")
	fmt.Scan(&filename)

	dat, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(dat), "\n")
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		lnameAndFname := strings.Split(line, " ")
		name := Name{fname: lnameAndFname[0], lname: lnameAndFname[1]}
		names = append(names, name)
	}

	for _, name := range names {
		fmt.Print(name.fname + " ")
		fmt.Print(name.lname + "\n")
	}
}
