package main

import (
	"fmt"
	"math"
)

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println(circleArea(1))
	cycleNames(names, greet)
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func greet(name string) {
	fmt.Println("Hello", name)
	return
}

func cycleNames(names []string, f func(string)) {
	for _, value := range names {
		f(value)
	}
	return
}
