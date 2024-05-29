// Given an array of positive integers nums and a positive integer target, return the minimal length of a
// subarray
//  whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Example 1:

// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.
// Example 2:

// Input: target = 4, nums = [1,4,4]
// Output: 1
// Example 3:

// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0

package main

import (
	"fmt"
	"math"
)

func main() {
	data := []int{2, 3, 1, 6, 4, 1}
	target := 7
	fmt.Println("min size of subarray is:", minSizeSubArray(data, target))
}

func minSizeSubArray(data []int, target int) int {
	low := 0
	high := 0
	minWindowSize := math.MaxInt
	sum := 0
	for high < len(data) {
		sum += data[high]
		high++
		//trying to reduce the window size
		for sum >= target {
			currentWindowSize := high - low
			minWindowSize = minVal(minWindowSize, currentWindowSize)
			sum -= data[low] // to reduce the sum starting from low pointer
			low++
		}
	}
	if minWindowSize == math.MaxInt {
		return 0
	}
	fmt.Println(low, high)
	return minWindowSize
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
