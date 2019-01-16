package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Word Finder
//
//   Your goal is to search for the words inside the corpus.
//
//   Note: This exercise is similar to the previous word finder program:
//   https://github.com/inancgumus/learngo/tree/master/13-loops/10-word-finder-labeled-switch
//
//   1. Get the search query from the command-line (it can be multiple words)
//
//   2. Filter these words, do not search for them:
//      and, or, was, the, since, very
//
//      To do this, use an array for the filtered words.
//
//   3. Print the words found.
//
// RESTRICTION
//   + The search and the filtering should be case insensitive
//
// HINT
//   + strings.Fields function converts a given string to a slice.
//
//     You can find its example in the word finder program that I've mentioned
//     above.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me a word to search.
//
//   go run main.go and was
//
//   go run main.go AND WAS
//
//   go run main.go cat beginning
//     #2 : "cat"
//     #11: "beginning"
//
//   go run main.go Cat Beginning
//     #2 : "cat"
//     #11: "beginning"
// ---------------------------------------------------------

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	words := strings.Fields(strings.ToLower(corpus))
	filter := [...]string{"and", "or", "was", "the", "since", "very"}
	var searchWords []string

	if len(os.Args) < 2 {
		fmt.Println("Please give me a word to search.")
		return
	}

	for _, a := range os.Args[1:] {
		var filtered bool
		for _, f := range filter {
			if a == f {
				filtered = true
			}
		}
		if !filtered {
			searchWords = append(searchWords, strings.ToLower(a))
		}
	}

	for _, s := range searchWords {

		for j, w := range words {
			if m, _ := regexp.MatchString(s, w); m {
				fmt.Printf("#%-2d : %q\n", j+1, w)
			}
		}
	}
}
