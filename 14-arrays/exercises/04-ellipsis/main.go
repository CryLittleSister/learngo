package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {
	var (
		names     = [...]string{"Cat", "Charlotte", "Alex"}
		distances = [...]int{40, 70, 120, 140, 200}
		data      = [...]byte{1, 14, 11, 91, 19}
		ratios    = [...]float64{9.99}
		alives    = [...]bool{true, true, false, true}
		zero      [0]byte
	)

	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}
	fmt.Println("")
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}
	fmt.Println("")
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}
	fmt.Println("")
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}
	fmt.Println("")
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}
	fmt.Println("")
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

	fmt.Println("-------------------------------------")
	fmt.Println("")

	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}
	fmt.Println("")
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}
	fmt.Println("")
	for i, v := range data {
		fmt.Printf("data[%d]: %d\n", i, v)
	}
	fmt.Println("")
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, v)
	}
	fmt.Println("")
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}
	fmt.Println("")
	for i, v := range zero {
		fmt.Printf("zero[%d]: %d\n", i, v)
	}
}
