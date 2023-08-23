package main

import "fmt"

func main() {
	menu := map[string]float64{ //maps in go have fixed types
		"Manu":  1.72,
		"Pedro": 1.81,
		"Joao":  1.90,
	}

	fmt.Println(menu)
	fmt.Println(menu["Manu"])

	// loop map
	for k, v := range menu {
		fmt.Println(k, v)
	}

	// ints as key type
	phonebook := map[int]string{
		123456789: "Manu",
		987654321: "Pedro",
	}
	phonebook[123456789] = "Joao" //updating value in map
	phonebook[4] = "Maria"        //adding value in map
	fmt.Println(phonebook)
}
