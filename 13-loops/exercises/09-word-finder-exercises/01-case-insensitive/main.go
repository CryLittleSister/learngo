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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------

func main() {
	const corpus = "You are a captive audience while sitting on the toilet, pet me. Feed me, give me attention. Scamper, rub face on everything"
	reg, _ := regexp.Compile(`[\.,]`)

	words := strings.Fields(strings.ToLower(reg.ReplaceAllString(corpus, "")))
	queries := os.Args[1:]
	if len(queries) < 1 {
		fmt.Println("Please provide at least one search term.")
	}
	var found bool
	for _, v := range queries {
		for i, w := range words {
			if w == strings.ToLower(v) {
				fmt.Printf("%q found at index #%d!!\n", v, i)
				found = true
				break
			}
		}
	}

	if !found {
		fmt.Println("Sorry, there are no results for any of your search terms")
	}

}
