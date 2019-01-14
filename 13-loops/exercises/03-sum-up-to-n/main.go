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
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func main() {
	args := os.Args
	var sum int

	if len(args) < 2 {
		fmt.Printf("ERROR: please provide 2 integers\n")
		return
	}
	min, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("ERROR: %q is not a number\n", args[1])
		return
	}
	max, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("ERROR: %q is not a number\n", args[2])
		return
	}
	if min >= max {
		fmt.Printf("ERROR: minimum number needs to be less than maximum\n")
		return
	}

	for i := min; i <= max; i++ {
		if i == max {
			sum += i
			fmt.Printf("%d = %d\n", i, sum)
		} else {
			sum += i
			fmt.Printf("%d + ", i)
		}
	}
}
