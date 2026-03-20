package main

import "fmt"

// jwtToken := 300000
var jwtToken = 300000

const LoginToken string = "weri-dfjs-jldf" // Public Variable

func main() {
	var username string = "Jatin"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.4554242523529308420
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherVariable2 string
	fmt.Println(anotherVariable2)
	fmt.Printf("Variable is of type: %T \n", anotherVariable2)

	// implicit type
	var website = "jatinporiya.netlify.app"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)
	// website = 3

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(jwtToken)
	fmt.Println(LoginToken)

}
