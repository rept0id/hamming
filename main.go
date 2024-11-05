package main

import (
	"fmt"
	"time"
)

/*** * * ***/

// factorizeUtil takes an integer and returns a slice of its factors
func factorizeUtil(n int) []int {
	var factors []int

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)

			n /= i
		}
	}

	return factors
}

// measureTimeUtil is a higher-order function that measures the time it takes to run any function.
func measureTimeUtil(fn func()) time.Duration {
	start := time.Now()

	fn() // Call the function

	return time.Since(start) // Return the elapsed time
}

/*** * * ***/

func generateHammingNums(x int) []int {
	hammingNumbers := []int{1} // Start with the first Hamming number

	for len(hammingNumbers) < x {
		currHamming := hammingNumbers[len(hammingNumbers)-1]
		nextHamming := nextHammingNum(currHamming)

		hammingNumbers = append(hammingNumbers, nextHamming)
	}

	return hammingNumbers
}

// nextHammingNum finds the next Hamming number
func nextHammingNum(x int) int {
	for {
		x++

		if isHammingNum(x) {
			return x
		}
	}
}

// nextNotHammingNum finds the next NOT Hamming number
func nextNotHammingNum(x int) int {
	for {
		x++

		if !isHammingNum(x) {
			return x
		}
	}
}

// isHammingNum checks if a number is a Hamming number
func isHammingNum(n int) bool {
	for n%2 == 0 {
		n /= 2
	}

	for n%3 == 0 {
		n /= 3
	}

	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}

/*** * * ***/

func main() {
	idx := 1000
	nums := generateHammingNums(idx)

	/*** * * ***/

	l := nums[idx-1] // last

	nextNotL := nextNotHammingNum(l) // next not Hamming of last

	tL := measureTimeUtil(func() { factorizeUtil(l) })               // time to factorize l
	tNextNotL := measureTimeUtil(func() { factorizeUtil(nextNotL) }) // time to factorize nextNotL

	/*** * * ***/

	fmt.Printf("Time taken to factorize %d: %v\n", l, tL)
	fmt.Printf("Time taken to factorize %d: %v\n", nextNotL, tNextNotL)
}
