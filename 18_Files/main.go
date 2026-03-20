package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in Golang")

	content := "This needs to go in a file - Jatin Poriya"
	file, err := os.Create("./myFile.txt")
	// if err != nil {
	// panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./myFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", databyte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
