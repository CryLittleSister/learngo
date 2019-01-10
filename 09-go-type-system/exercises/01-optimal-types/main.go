// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"math"
)

// ---------------------------------------------------------
// EXERCISE: Optimal Types
//
//  1. Choose the optimal data types for the given situations.
//  2. Print them all
//  3. Try converting them to lesser data types.
//     For example try converting int64 variable to int32.
//     Then observe the result.
//     Search the web why the result is so?
//
// NOTE
//  This is just an exercise for teaching you different
//  data types. Do not apply it to the real-life.
//
//  As I said in the lectures that, premature optimization
//  is not a good thing.
// ---------------------------------------------------------

func main() {
	// DONT FORGET: There are also unsigned data types.
	//              (For positive numbers)

	// DO NOT USE: int data type
	// Use only the ones with the bitsizes

	// ---

	// an english letter (search web for: ascii control code)
	var el byte = 'A'
	fmt.Println(el)

	// a non-english letter (search web for: unicode codepoint)
	var nel int16 = 'Ñ'
	fmt.Println(nel)

	// a year in 4 digits like 2040
	var y uint16 = 9999
	fmt.Println(y)

	// a month in 2 digits: 1 to 12
	var m uint8 = 10
	fmt.Println(m)

	// the speed of the light
	var sol uint32 = 299792458
	fmt.Println(sol)

	// angle of a circle
	var ac float32 = 345
	fmt.Println(ac)

	// PI number: 3.141592653589793
	var pi float64 = math.Pi
	fmt.Println(pi)

}
