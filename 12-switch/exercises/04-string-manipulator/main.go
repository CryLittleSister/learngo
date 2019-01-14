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
	"strings"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate the
//  given strings to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

func main() {

	msg, args := `[command] [string]
		
Available commands: lower, upper and title`, os.Args

	if len(args) < 3 {
		fmt.Println(msg)
		return
	}

	com, str := args[1], args[2]

	switch com {
	case "lower":
		msg = strings.ToLower(str)
	case "upper":
		msg = strings.ToUpper(str)
	case "title":
		msg = strings.Title(str) + " xxxxx"
	default:
		msg = "Unknown command: \"" + com + "\""
	}

	fmt.Println(msg)
}
