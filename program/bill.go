package main

import "fmt"

type bill struct { //class
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 1.02, "coke": 2.1},
		tip:   0,
	}
	return b
}

// receiver function (method)
func (b bill) format() string { //this associates bill struct with method format()
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v %v \n", k+":", v) //-25 means spaces to the right
		total += v
	}

	fs += fmt.Sprintf("%-25v %v", "Total:", total)

	return fs
}
