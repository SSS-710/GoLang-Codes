package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	tags     []string
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 200, "LearnCodeOnline.in", "abc123", []string{"web-dev"}},
		{"full stack Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev"}},
		{"django Bootcamp", 259, "LearnCodeOnline.in", "abc123", []string{"web-dev"}},
	}
	// package this data as JSON

	finalJson, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
