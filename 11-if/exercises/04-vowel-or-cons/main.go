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
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a letter")
		return
	}

	var l string

	if len(os.Args[1]) > 1 {
		fmt.Println("Give me a letter")
		return
	} else {
		l = os.Args[1]
	}

	if strings.IndexAny(l, "aeiouwy") == -1 {
		fmt.Printf("%q is a consonant.\n", l)
	} else if strings.IndexAny(l, "aeiou") == -1 {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", l)
	} else {
		fmt.Printf("%q is a vowel.\n", l)
	}
}
