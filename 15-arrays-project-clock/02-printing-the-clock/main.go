// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	t := time.Now()
	hour, min, sec := t.Hour(), t.Minute(), t.Second()
	h0, _ := strconv.Atoi(strings.Split(strconv.Itoa(hour), "")[0])
	h1, _ := strconv.Atoi(strings.Split(strconv.Itoa(hour), "")[1])
	m0, _ := strconv.Atoi(strings.Split(strconv.Itoa(min), "")[0])
	m1, _ := strconv.Atoi(strings.Split(strconv.Itoa(min), "")[1])
	s0, _ := strconv.Atoi(strings.Split(strconv.Itoa(sec), "")[0])
	s1, _ := strconv.Atoi(strings.Split(strconv.Itoa(sec), "")[1])

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	clock := [...]placeholder{
		digits[h0], digits[h1],
		colon,
		digits[m0], digits[m1],
		colon,
		digits[s0], digits[s1],
	}

	for line := range clock[0] {
		for digit := range clock {
			fmt.Print(clock[digit][line], "  ")
		}
		fmt.Println()
	}
}
