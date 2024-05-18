package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func takeInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := takeInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := takeInput("Item name: ", reader)
		price, _ := takeInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number...")
			promptOptions(b)
		}
		b.updateItems(name, p)

		fmt.Println("item added -", name, price)
		promptOptions(b)

	case "t":
		tip, _ := takeInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number...")
			promptOptions(b)
		}
		b.addTip(t)

		fmt.Println("tip has been updated to", tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("bill has been saved as", b.name)

	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := takeInput("Creating bill for?", reader)
	b := newBill(name)
	return b
}

func main() {
	theBill := createBill()
	promptOptions(theBill)
}
