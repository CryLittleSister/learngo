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
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  Player wins: then randomly printone of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  Player loses: then randomly printone of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
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

	var win, lose string

	for i := 1; i <= maxTurns; i++ {
		switch rand.Intn(3) {
		case 0:
			win = "%d!! WEEEE ARE THE CHAMPIONS MY FRIEEEEND\n"
			lose = "BOOOOOOOOOO"
		case 1:
			win = "%d is correct!! WELL DONE!!!\n"
			lose = "Sorry, you lose. Try again?"
		case 2:
			win = "You're right! %d is spot on, way to go champ!\n"
			lose = "uh-oh you looooose sucka"
		}

		var cpu int
		if guess < 10 {
			cpu = rand.Intn(11)
		} else {
			cpu = rand.Intn(guess + 1)
		}
		if guess == cpu {
			fmt.Printf(win, guess)
			return
		} else if i == maxTurns {
			fmt.Println(lose)
			return
		}

	}
}
