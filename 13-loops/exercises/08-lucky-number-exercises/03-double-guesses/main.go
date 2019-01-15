// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	maxTurns := 5
	intro := `Welcome to the Lucky Number Game! 🍀

The program will pick %d random positive numbers.
Your mission is to guess one of those numbers with two guesses.
	
The greater your number is, harder it gets.
	
Wanna play?
`
	if len(args) != 2 {
		fmt.Printf(intro, maxTurns)
		return
	}

	guess1, err1 := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil || guess1 < 0 || guess2 < 0 {
		fmt.Println("ERROR: Please choose 2 positive numbers")
		return
	}

	if guess2 > guess1 {
		guess1, guess2 = guess2, guess1
	}

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= maxTurns; i++ {
		var cpu int
		if guess1 < 10 {
			cpu = rand.Intn(11)
		} else {
			cpu = rand.Intn(guess1 + 1)
		}

		if guess1 == cpu || guess2 == cpu {
			fmt.Println("WEEEE ARE THE CHAMPIONS MY FRIEEEEND")
			return
		} else if i == maxTurns {
			fmt.Println("uh-oh you looooose sucka")
			return
		}

	}
}
