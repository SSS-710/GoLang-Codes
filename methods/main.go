package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	// no inheritence i golang; No super or parent

	sss := User{"sss", "sss@go.dev", true, 19}
	fmt.Println(sss)
	fmt.Printf("sss details are: %+v\n", sss)
	sss.GetStatus()
	sss.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)

}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
