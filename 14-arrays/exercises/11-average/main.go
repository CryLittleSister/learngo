// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

func main() {
	if len(os.Args) < 2 || len(os.Args) > 6 {
		fmt.Println("Please give me up to 5 numbers")
		return
	}

	var nums [5]float64
	var sum float64
	var numOfNums float64
	for i, n := range os.Args[1:] {
		if n, err := strconv.ParseFloat(n, 64); err == nil {
			nums[i] = n
			sum += n
			numOfNums++
		}
	}

	fmt.Printf("Your numbers: %v\n", nums)
	fmt.Printf("Average: %.2f\n", sum/numOfNums)
}
