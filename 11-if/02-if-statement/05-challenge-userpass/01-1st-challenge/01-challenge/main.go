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
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	var u, p string
	if len(os.Args) >= 3 {
		u, p = os.Args[1], os.Args[2]
	} else {
		fmt.Println("Usage: [username] [password]")
		return
	}

	if u == "jack" {
		if p == "1888" {
			fmt.Println("Access granted to \"jack\".")
		} else {
			fmt.Println("Invalid password for \"jack\".")
		}
	} else {
		fmt.Printf("Access denied for %q.\n", u)
	}
}
