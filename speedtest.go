package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func generateDivisibleNumbers() (int, int) {
	rand.Seed(time.Now().UnixNano())

	// Generate a random three-digit divisor between 2 and 9
	divisor := generateRandomNumber(2, 9)

	// Generate a random dividend that is divisible by the divisor
	dividend := rand.Intn(999/divisor+1) * divisor
	return dividend, divisor
}

func main() {
	// Perform 5 iterations
	for i := 0; i < 5; i++ {
		// Generate expression for sum
		a := generateRandomNumber(100, 999)
		b := generateRandomNumber(10, 999)
		// Print the expression
		fmt.Printf("%d + %d = ", a, b)
		fmt.Println()

		// Generate expression for subtraction
		c := generateRandomNumber(100, 999)
		d := generateRandomNumber(10, c)
		// Print the expression
		fmt.Printf("%d - %d = ", c, d)
		fmt.Println()

		// Generate expression for multiplication
		e := generateRandomNumber(100, 999)
		f := generateRandomNumber(2, 9)
		// Print the expression
		fmt.Printf("%d x %d = ", e, f)
		fmt.Println()

		// Generate expression for division
		g, h := generateDivisibleNumbers()
		// Print the expression
		fmt.Printf("%d \u00F7 %d = ", g, h)

		// You can perform any additional calculations or operations here if needed
		// For simplicity, this example doesn't include the actual result calculation

		fmt.Println()
	}
}
