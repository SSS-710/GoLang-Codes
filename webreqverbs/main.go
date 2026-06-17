package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("creating webverb video")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {

	const myurl = "http://localhost:3000"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("content length is:", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(content)
	//fmt.Println(string(content))

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	//fake json payload

	requestBody := strings.NewReader(`

		{
			"coursename": "Lets go with golang",
			"price" : 0,
			"platform": "learnCodeOnline.in"

		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "sss")
	data.Add("lastname", "yoyo")
	data.Add("email", "sss@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
