package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the contants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go invalid
//     Invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 9.24 EUR
//     10.50 USD is 8.19 GBP
//     10.50 USD is 1186.71 JPY
//
//   go run main.go 1
//     1.00 USD is 0.88 EUR
//     1.00 USD is 0.78 GBP
//     1.00 USD is 113.02 JPY
// ---------------------------------------------------------

func main() {
	const (
		EUR = iota
		GBP
		JPY
	)

	conv := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}
	if len(os.Args) < 2 {
		fmt.Println("Please enter an amount in dollars to convert.")
		return
	}

	USD, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("ERROR: %s is not a number.\n", os.Args[1])
		return
	}

	for i, v := range conv {
		var cur string
		switch i {
		case 0:
			cur = "EUR"
		case 1:
			cur = "GBP"
		case 2:
			cur = "JPY"
		}
		fmt.Printf("%.2f USD is %.2f %v\n", USD, v*USD, cur)
	}
}
