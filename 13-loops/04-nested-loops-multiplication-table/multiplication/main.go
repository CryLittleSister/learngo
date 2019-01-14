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

// EXERCISE: Get the `max` from the user
//           And create the table according to that.

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please enter an integer")
		return
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("ERROR: %q is not an integer\n", args[1])
	}

	// print the header
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
