package main

import (
	"fmt"
)

// passing a pointer would allow us to change the value of the variable
func updateName(x *string) { // *string is a pointer to a string
	*x = "wedge" // *x is the value at the memory address
}

func main() {
	name := "John"

	// fmt.Println("memory address of name variable is", &name)

	m := &name // m is a pointer to the memory address of name
	fmt.Println("memory address of m variable is", m)
	fmt.Println("value at memory address", *m) // *m is the value at the memory address

	updateName(m)
	fmt.Println(name)

	// fmt.Println(name)
}
