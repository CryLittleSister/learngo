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

// NOTE: You should always pass it at least one argument

func main() {
	msg := os.Args[1]
	l := len(msg)
	r := strings.Repeat("!", l)

	s := r + msg + r
	s = strings.ToUpper(s)

	fmt.Println(s)
}
