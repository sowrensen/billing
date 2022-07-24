package main

import (
	"fmt"
	"os"
)

type voucher struct {
	name  string
	items map[string]float64
	tip   float64
}

func newVoucher(name string) voucher {
	v := voucher{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return v
}

func (v *voucher) format() string {
	fs := fmt.Sprintf("Voucher for %v \n", v.name)
	var total float64 = 0

	for menu, price := range v.items {
		fs += fmt.Sprintf("%-25v ...$%0.2f \n", menu+":", price)
		total += price
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Tip:", v.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Total:", total + v.tip)

	return fs
}

func (v *voucher) addItem(name string, price float64) {
	v.items[name] = price
}

func (v *voucher) addTip(tip float64) {
	fmt.Println("Entered tip: ", tip)
	v.tip = tip
}

func (v *voucher) save() {
	data := []byte(v.format())

	path := "vouchers/" + "Voucher of " + v.name + ".txt"
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Voucher has been saved.")
}
