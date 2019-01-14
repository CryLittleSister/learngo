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
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	args := os.Args

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
	var (
		sum int
		i   = min - 1
	)
	for {
		i++
		if i%2 != 0 && i != max {
			continue
		}
		if i > max {
			break
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
