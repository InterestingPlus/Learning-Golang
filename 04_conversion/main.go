package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Printf("Please rate our pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n') // String
	fmt.Println("Thanks for Rating,", input)
	fmt.Printf("Type of this Rating is %T \n", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // Number

	if err != nil {
		fmt.Println(err)
		// panic(err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}

}
