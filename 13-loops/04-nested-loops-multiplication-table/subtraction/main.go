package main

import (
	"fmt"
	"os"
	"strconv"
)

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
	fmt.Printf("%5s", "-")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i-j)
		}
		fmt.Println()
	}
}
