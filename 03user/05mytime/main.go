package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to timr study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.Local)
	fmt.Println("createdDate")
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

// new () MEMORY MANAGEMENT = allocate memo but no init = zeroed storage
// make() allocate memo and init  = non-zeroed storage
