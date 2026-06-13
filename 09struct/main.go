package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	// no inheritence i golang; No super or parent

	sss := User{"sss", "sss@go.dev", true, 19}
	fmt.Println(sss)
	fmt.Printf("sss details are: %+v\n", sss)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
