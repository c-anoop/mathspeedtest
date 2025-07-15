package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var rng *rand.Rand

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func generateRandomNumber(min, max int) int {
	return rng.Intn(max-min+1) + min
}

func generateDivisibleNumbers() (int, int) {
	// Generate a random three-digit divisor between 2 and 9
	divisor := generateRandomNumber(2, 9)

	// Generate a random dividend that is divisible by the divisor
	dividend := rng.Intn(999/divisor+1) * divisor
	return dividend, divisor
}

func generateTableDodging() (int, int) {
	// Generate two random numbers where one is from 1-10 (table)
	table := generateRandomNumber(1, 10)
	multiplier := generateRandomNumber(1, 10)

	return table, multiplier
}

func generateSpeedTestProblem(problemType int) {
	switch problemType {
	case 0: // Addition
		a := generateRandomNumber(100, 999)
		b := generateRandomNumber(10, 999)
		fmt.Printf("%d + %d = \n", a, b)
	case 1: // Subtraction
		c := generateRandomNumber(100, 999)
		d := generateRandomNumber(10, c)
		fmt.Printf("%d - %d = \n", c, d)
	case 2: // Multiplication
		e := generateRandomNumber(100, 999)
		f := generateRandomNumber(2, 9)
		fmt.Printf("%d x %d = \n", e, f)
	case 3: // Division
		g, h := generateDivisibleNumbers()
		fmt.Printf("%d ÷ %d = \n", g, h)
	}
}

func generateTableRow() {
	table, multiplier := generateTableDodging()
	fmt.Printf("%d x %d = \n", table, multiplier)
}

func main() {
	// Check if we have the required arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: speedtest <operation> <count>")
		fmt.Println("  operation: 'speed' or 'table'")
		fmt.Println("  count: number of rows to generate (integer)")
		os.Exit(1)
	}

	operation := os.Args[1]
	countStr := os.Args[2]

	// Parse count argument
	count, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid integer\n", countStr)
		os.Exit(1)
	}

	if count <= 0 {
		fmt.Println("Error: count must be a positive integer")
		os.Exit(1)
	}

	// Validate operation
	if operation != "speed" && operation != "table" {
		fmt.Printf("Error: operation must be 'speed' or 'table', got '%s'\n", operation)
		os.Exit(1)
	}

	// Generate the requested test
	fmt.Printf("Math Speed Test - %s mode (%d problems)\n", operation, count)
	fmt.Println("========================================")

	for i := 0; i < count; i++ {
		if operation == "speed" {
			// Cycle through all 4 operation types for variety
			generateSpeedTestProblem(i % 4)
		} else {
			generateTableRow()
		}
	}
}
