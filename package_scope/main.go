package main

import "fmt"

// to access var & func from another file,
// can run it in terminal same time as this file

var score = 95.5

func main() {
	sayHello("John")

	for _, v := range points {
		fmt.Print(v, " ")
	}

	showScore()

}
