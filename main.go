package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	// read and trim input
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err
}

func createVoucher() voucher {
	// create a new io buffer
	reader := bufio.NewReader(os.Stdin)
	name, _ := readInput("What's the customer name?: ", reader)

	// create a new voucher
	voucher := newVoucher(name)
	fmt.Println("Created voucher for:", name)
	return voucher
}

func promptOptions(v voucher) {
	// create a new io buffer
	reader := bufio.NewReader(os.Stdin)
	choice, _ := readInput("a - Add item, t - Add tip, v - View voucher, s - Save voucher, e - Exit: ", reader)

	switch choice {
	case "a":
		addItem(v, reader)
	case "t":
		addTip(&v, reader)
	case "v":
		fmt.Println(v.format())
	case "s":
		v.save()
	case "e":
		fmt.Println("Closing program...")
		os.Exit(0)
	default:
		fmt.Println("Not a valid choice, please try again.")
	}

	promptOptions(v)
}

func addItem(v voucher, reader *bufio.Reader) {
	name, _ := readInput("Menu name: ", reader)
	price, _ := readInput("Menu price: ", reader)
	p, err := strconv.ParseFloat(price, 64)

	if err != nil {
		fmt.Println("Price must be a number!")
		return
	}

	v.addItem(name, p)
	fmt.Printf("Item added - %v, $%v \n", name, p)
}

func addTip (v *voucher, reader *bufio.Reader) {
	tip, _ := readInput("Tip amount: ", reader)
	t, err := strconv.ParseFloat(tip, 64)

	if err != nil {
		fmt.Println("Amount must be a number!")
		return
	}

	v.addTip(t)
	fmt.Printf("Tip added - $%v \n", t)
}

func main() {
	myVoucher := createVoucher()
	promptOptions(myVoucher)
}
