package main

import (
	"fmt"
	"os"
)

type bill struct { //class
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
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

	fs += fmt.Sprintf("%-25v %v \n", "Tip:", b.tip)
	total += b.tip

	fs += fmt.Sprintf("%-25v %v", "Total:", total)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) { //use pointer to update the value
	b.tip = tip //b is automatically dereferenced
}

// add item to bill
func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) saveBill() {
	//write to file
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644) //0644 is the permission

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
