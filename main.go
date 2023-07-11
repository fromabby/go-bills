package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		fPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, fPrice)
		promptOptions(b)
		break
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		fTip, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(fTip)
		promptOptions(b)
		break
	case "s":
		b.save()
		break
	default:
		fmt.Println("not a valid option...")
		promptOptions(b)
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {
	mybill := createBill()

	promptOptions(mybill)
}
