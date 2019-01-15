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
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game to easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	maxTurns := 5
	intro := `Welcome to the Lucky Number Game! 🍀

The program will pick a selection of random positive numbers based on how high your guess is.
Your mission is to guess one of those numbers.
	
The greater your number is, harder it gets.
	
Wanna play?`
	if len(args) < 1 {
		fmt.Println(intro)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil || guess < 0 {
		fmt.Println("ERROR: Please choose a positive number")
		return
	}

	switch {
	case guess >= 100:
		maxTurns += 20
	case guess >= 50:
		maxTurns += 10
	case guess >= 25:
		maxTurns += 5
	case guess >= 10:
		maxTurns++
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
