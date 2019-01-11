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
	"regexp"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
	} else {
		num, _ := strconv.ParseInt(os.Args[1], 10, 64)
		m, _ := regexp.MatchString("[0-9]", os.Args[1])
		if m {
			if num%2 == 0 {
				if num%8 == 0 {
					fmt.Printf("%d is an even number and it's divisible by 8\n", num)
				} else {
					fmt.Printf("%d is an even number\n", num)
				}
			} else {
				fmt.Printf("%d is an odd number\n", num)
			}
		} else {
			fmt.Printf("%q is not a number\n", os.Args[1])
		}
	}

}
