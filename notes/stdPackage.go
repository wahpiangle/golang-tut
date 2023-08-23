// standard library
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Strings package
	/*
		greeting := "Hello there friends!"
		fmt.Println(strings.Contains(greeting, "ll")) //true
		fmt.Println(strings.ToUpper(greeting))
		fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

		fmt.Println(strings.Index(greeting, "ll")) //shows the index of the first instance of "ll"

		fmt.Println(strings.Split(greeting, " ")) //splits the string into a slice

	*/

	// sort
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages) //this will change the slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) //returns the index of the value 30
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Print(sort.SearchStrings(names, "bowser")) //returns the index of the value "bowser"
}
