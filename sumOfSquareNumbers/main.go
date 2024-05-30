// Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.

// Example 1:

// Input: c = 5
// Output: true
// Explanation: 1 * 1 + 2 * 2 = 5
// Example 2:

// Input: c = 3
// Output: false

package main

import (
	"fmt"
	"math"
)

func main() {
	c := 2
	fmt.Println("is it sum of square numbers:", checkSumOfSquareNumber(c))
}

func checkSumOfSquareNumber(c int) bool {
	temp := math.Sqrt(float64(c))
	data := int(temp)
	low := 0
	high := data
	fmt.Println(high)
	sum := 0
	for low <= data {
		sum = int(math.Pow(float64(high), 2)) + int(math.Pow(float64(low), 2))
		if sum < c {
			low++
		} else if sum > c {
			high--
		}
		if sum == c {
			return true
		}
	}
	return false
}
