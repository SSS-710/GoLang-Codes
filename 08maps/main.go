package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	// loops are interset in golang

	for _, value := range languages {
		fmt.Println("For key v, value is v\n", value)
	}
}
