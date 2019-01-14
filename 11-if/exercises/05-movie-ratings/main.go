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
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

func main() {
	var (
		age float64
		err error
	)
	if len(os.Args) > 1 {
		age, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Println("ERROR: Please give an age")
		}

		if os.Args[1] != "0" && age != 0 {
			if age < 1 {
				fmt.Printf("Wrong age: %q\n", os.Args[1])
			} else if age <= 12 {
				fmt.Println("PG-Rated")
			} else if age > 12 && age <= 17 {
				fmt.Println("PG-13")
			} else {
				fmt.Println("R-Rated")
			}

		} else {
			fmt.Println("Requires age")
		}
	} else {
		fmt.Println("Requires age")
	}
}
