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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {
	var turn int
	args := os.Args[1:]
	maxTurns := 5
	intro := `Welcome to the Lucky Number Game! 🍀

The program will pick %d random positive numbers.
Your mission is to guess one of those numbers.
	
The greater your number is, harder it gets.
	
Wanna play?
`
	if len(args) != 1 {
		fmt.Printf(intro, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil || guess < 0 {
		fmt.Printf("ERROR: %q is not a valid positive number\n", args[0])
		return
	}

	rand.Seed(time.Now().UnixNano())

	for {
		var cpu int
		if guess < 10 {
			cpu = rand.Intn(11)
		} else {
			cpu = rand.Intn(guess + 1)
		}
		turn++
		if guess == cpu && turn == 1 {
			fmt.Println("OMG FIRST GO!!!!")
			return
		} else if guess == cpu && turn <= maxTurns {
			fmt.Printf("%d is correct!! WELL DONE!!!\n", guess)
			return
		} else if turn == maxTurns {
			fmt.Println("Sorry, you lose. Try again?")
			return
		}

	}

}
