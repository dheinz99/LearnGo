package main
import "fmt"
import "bufio"
import "os"
import "log"
import "strings"
type person struct{
	fname string
	lname string
}
func main() {
	var filename string
	fmt.Scan(&filename)
	slice:=[]person{}
	readFile, err := os.Open(filename)
 
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string
	
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
 
	readFile.Close()
 
	for _, eachline := range fileTextLines {
		s:=strings.Split(eachline," ")
		p1:=person{s[0],s[1]}
		slice=append(slice,p1)
	}
	fmt.Println(slice)
}