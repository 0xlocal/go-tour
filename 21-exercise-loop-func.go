// As a way to play with functions and loops, let's implement a square root function: given a number x,
// we want to find the number z for which z² is most nearly x.
//
// Computers typically compute the square root of x using a loop. Starting with some guess z,
// we can adjust z based on how close z² is to x, producing a better guess:
//
// z -= (z*z - x) / (2*z)
// Repeating this adjustment makes the guess better and better until we reach an answer that is as close
// to the actual square root as can be.
//
// Implement this in the func Sqrt provided. A decent starting guess for z is 1, no matter what the input.
// To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to
// the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.
//
// Hint: To declare and initialize a floating point value, give it floating point syntax or use a conversion:
//
// z := 1.0
// z := float64(1)
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var t float64

	// solution
	for {
		z, t = z-(z*z-x)/(2*z), z

		if math.Abs(t-z) < 1e-6 {
			break
		}
	}

	// former test
	// for i := 0; i < 10; i++ {
	// 	z -= (z*z - x) / (2 * z)
	// 	fmt.Printf("Iteration %v, value = %v\n", i, z)
	// }

	return z
}

func main() {
	guess := Sqrt(2)
	expected := math.Sqrt(2)

	fmt.Printf("Guess: %v, Expected: %v, Error: %v", guess, expected, math.Abs(guess-expected))
}
