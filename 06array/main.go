package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list is:", len(fruitList))

	var vegList = [3]string{"potato", "ladyfinger", "ghuiya"}
	fmt.Println("vegy list is:", vegList)
}
