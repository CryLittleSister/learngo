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
// EXERCISE: Enough Picks
//
//  If the player's guessnumber is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guessnumber is above 10; then the
//  computer should generate the numbers
//  between 0 and the guessnumber.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guessnumber is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	maxTurns := 5
	intro := `Welcome to the Lucky Number Game! 🍀

The program will pick %d random positive numbers.
Your mission is to guess one of those numbers.
	
The greater your number is, harder it gets.
	
Wanna play?
`
	if len(args) < 1 {
		fmt.Printf(intro, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil || guess < 0 {
		fmt.Println("ERROR: Please choose a positive number")
		return
	}

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= maxTurns; i++ {
		var cpu int
		if guess < 10 {
			cpu = rand.Intn(11)
		} else {
			cpu = rand.Intn(guess + 1)
		}

		if guess == cpu {
			fmt.Println("WEEEE ARE THE CHAMPIONS MY FRIEEEEND")
			return
		} else if i == maxTurns {
			fmt.Println("uh-oh you looooose sucka")
			return
		}

	}
}
