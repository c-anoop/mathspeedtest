/*
MIT License

# Copyright (c) 2023 Anoop Chargotra

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
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

func generateSpeedTestProblem(problemType int) string {
	switch problemType {
	case 0: // Addition
		a := generateRandomNumber(100, 999)
		b := generateRandomNumber(10, 999)
		return fmt.Sprintf("%d + %d = ", a, b)
	case 1: // Subtraction
		c := generateRandomNumber(100, 999)
		d := generateRandomNumber(10, c)
		return fmt.Sprintf("%d - %d = ", c, d)
	case 2: // Multiplication
		e := generateRandomNumber(100, 999)
		f := generateRandomNumber(2, 9)
		return fmt.Sprintf("%d x %d = ", e, f)
	case 3: // Division
		g, h := generateDivisibleNumbers()
		return fmt.Sprintf("%d ÷ %d = ", g, h)
	}
	return ""
}

func generateTableProblem() string {
	table, multiplier := generateTableDodging()
	return fmt.Sprintf("%d x %d = ", table, multiplier)
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

	if count > 50 {
		fmt.Println("Error: count cannot be more than 50")
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

	for i := 0; i < count; i += 2 {
		var problem1, problem2 string

		if operation == "speed" {
			// Cycle through all 4 operation types for variety
			problem1 = generateSpeedTestProblem(i % 4)
			if i+1 < count {
				problem2 = generateSpeedTestProblem((i + 1) % 4)
			}
		} else {
			problem1 = generateTableProblem()
			if i+1 < count {
				problem2 = generateTableProblem()
			}
		}

		// Print in two columns with proper spacing
		if problem2 != "" {
			fmt.Printf("%-25s        %s\n", problem1, problem2)
		} else {
			fmt.Printf("%s\n", problem1)
		}
	}
}
