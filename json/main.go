package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"_"`
	tags     []string
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 200, "LearnCodeOnline.in", "abc123", []string{"web-dev"}},
		{"full stack Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev"}},
		{"django Bootcamp", 259, "LearnCodeOnline.in", "abc123", []string{"web-dev"}},
	}
	// package this data as JSON

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	 {
                "coursename": "full stack Bootcamp",
                "Price": 299,
                "website": "LearnCodeOnline.in",
                "_": "abc123"
    }
	
	
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)

	} else {

		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("%#v\n", myOnlineData)

}
