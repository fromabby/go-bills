package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")

	mybill.addItem("clubhouse", 65)
	mybill.addItem("lechon kawali", 145)
	mybill.addItem("fries", 35)

	mybill.updateTip(5)

	fmt.Println(mybill.format())
}
