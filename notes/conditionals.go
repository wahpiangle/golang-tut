package main

import "fmt"

func main() {
	age := 45
	fmt.Println(age <= 50)

	if age < 30 {
		fmt.Println("You are young")
	} else if age < 40 {
		fmt.Println("You are getting there")
	} else {
		fmt.Println("You are old")
	}

	names := []string{"John", "Paul", "George", "Ringo"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break
		}

		fmt.Printf("The value at %v is %v\n", index, value)
	}
}
