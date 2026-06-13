package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("index is and value is %v\n", day)

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			continue
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++

	}

lco:
	fmt.Println("Jumping at LearnCodeonline.in")
}
