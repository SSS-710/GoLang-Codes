package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer myDefer()
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
}

// world, one , two,

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}
}
