package main

import (
	"fmt"
	"strings"
)

// this function returns 2 strings
func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string

	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func main() {
	first, sn := getInitials("John Doe")
	fmt.Println(first, sn)
}
