package main

import (
	"fmt"
)

const LoginToken string = "shashank" // public

func main() {
	var username string = "sss"
	fmt.Println("usrename")
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.5
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style

	numberOfUser := 30000.0
	fmt.Println(numberOfUser)

	fmt.Println("LoginToken")
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
