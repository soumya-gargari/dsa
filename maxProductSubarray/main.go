// Given an integer array nums, find a
// subarray
//  that has the largest product, and return the product.

// The test cases are generated so that the answer will fit in a 32-bit integer.

// Example 1:

// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.
// Example 2:

// Input: nums = [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

package main

import "fmt"

func main() {
	data := []int{2, 3, 0, -5, -2, 4}
	fmt.Println("max subarray product is:", maxProductSubArray(data))
}

// Time complexity : O(n)
// space complexity: O(1)
func maxProductSubArray(data []int) int {
	leftProduct, rightProduct := 1, 1
	ans := data[0]
	n := len(data)
	for i := 0; i < n; i++ {
		if leftProduct == 0 {
			leftProduct = 1
		}
		if rightProduct == 0 {
			rightProduct = 1
		}
		leftProduct *= data[i]
		rightProduct *= data[n-1-i]
		ans = Max(ans, Max(leftProduct, rightProduct))
		fmt.Println(leftProduct, rightProduct)
	}

	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
