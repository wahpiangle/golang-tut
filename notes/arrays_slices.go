package main

import "fmt"

func main() { //arrays are fixed length
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30} //shorthand

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi" //change value of array

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices
	var scores = []int{100, 50, 60} //no length
	scores = append(scores, 85)     //add to slice
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]  //start at 1, end at 3
	rangeTwo := names[2:]   //start at 2, end at end
	rangeThree := names[:3] //start at start, end at 3
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}
