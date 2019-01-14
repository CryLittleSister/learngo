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
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	args := os.Args
	var sum int

	if len(args) < 2 {
		fmt.Printf("ERROR: please provide 2 integers\n")
		return
	}
	min, err1 := strconv.Atoi(args[1])
	max, err2 := strconv.Atoi(args[2])
	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("ERROR: please enter valid integers where min < max")
		return
	}

	if max%2 != 0 {
		max--
	}

	for i := min; i <= max; i++ {
		if i%2 != 0 && i != max {
			continue
		}
		if i < max {
			fmt.Printf("%d + ", i)
			sum += i
		} else {
			sum += i
			fmt.Printf("%d = %d ", i, sum)
		}
	}

}
