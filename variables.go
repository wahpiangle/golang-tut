package main

import "fmt"

func main() {

	// strings
	var nameOne string = "mario" //typed
	var nameTwo = "luigi"        //no type
	var nameThree string         //no value
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi" //shorthand for var nameFour string = "yoshi"
	fmt.Println(nameFour)

	// ints
	var ageOne int64 = 2000000000000
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255 //unsigned int
	fmt.Println(numOne, numTwo, numThree)

	// floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 25.98
	scoreThree := 25.98
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	//print format
	fmt.Print("print statement \n")

	name := "AHH"
	age := 12
	fmt.Printf("Hello, %s! The number is %d\n", name, age)
}
