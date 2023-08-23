package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// strings, ints, bools, floats, arrays, structs, and bools are all passed by value
	name := "John"
	name = updateName(name)
	println(name)

	// slices, maps, and channels are passed by reference, so value will be updated
	menu := map[string]float64{
		"item1": 100,
		"item2": 200,
	}
	updateMenu(menu)
	fmt.Println(menu)

}
