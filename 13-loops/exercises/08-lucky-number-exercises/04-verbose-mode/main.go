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
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	maxTurns := 5
	intro := `Welcome to the Lucky Number Game! 🍀

The program will pick %d random positive numbers.
Your mission is to guess one of those numbers.
	
The greater your number is, harder it gets.
	
Wanna play?

(Provide -v flag to see the picked numbers.)
`
	if len(args) < 1 {
		fmt.Printf(intro, maxTurns)
		return
	}
	var flag int
	if args[0] == "-v" {
		flag = 1
	}

	guess, err := strconv.Atoi(args[flag])
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
		if flag == 1 {
			fmt.Printf("%d.. ", cpu)
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
