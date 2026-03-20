package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the Rating for our Pizza: ")

	// comma ok syntax || err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating,", input)
	fmt.Printf("Type of this Rating is %T \n", input)
}
