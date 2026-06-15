package main

import (
	"fmt"
	"io"

	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("Web request")

	http.Get(url)

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // to close the connections
	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
