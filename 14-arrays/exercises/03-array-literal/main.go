package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {
	var (
		names     = [3]string{"Cat", "Charlotte", "Alex"}
		distances = [5]int{40, 70, 120, 140, 200}
		data      = [5]byte{1, 14, 11, 91, 19}
		ratios    = [1]float64{9.99}
		alives    = [4]bool{true, true, false, true}
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
