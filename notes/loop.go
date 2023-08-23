package main

import (
	"fmt"
)

func main() {

	//! For loop
	// x := 0
	// for x < 5 {
	// 	fmt.Println("Value of x is ", x)
	// 	x++
	// }

	// ! Cycle through array
	names := []string{"John", "Doe", "Smith"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// ! Range for loop
	// for index, value := range names {
	// 	fmt.Printf("The value at index %v is %v \n", index, value)
	// }

	// ! Range for loop without index
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
		value = "New" // This will not change the value in the array
	}

	fmt.Println(names)
}
