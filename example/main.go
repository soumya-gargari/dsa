// Given an array A[] of n numbers and another number x, the task is to check whether or not there exist two elements in A[] whose sum is exactly x
// Input: arr[] = {0, -1, 2, -3, 1}, x= -2
// Output: [3,4] // index of the numbers used
// Explanation:  If we calculate the sum of the output,1 + (-3) = -2
// Input: arr[] = {1, -2, 1, 0, 5}, x = 0
// Output: []

package main

import (
	"fmt"
)

func findTwoSumIndices(arr []int, x int) (int, int, bool) {
	// Create a map to store the difference of x and each element
	hashMap := make(map[int]int)
	for i, num := range arr {
		if idx, ok := hashMap[num]; ok {
			return idx, i, true
		}
		hashMap[x-num] = i
	}
	// Return false if no pair is found
	return 0, 0, false
}

func main() {
	// Example usage
	arr := []int{10, 2, 3, 7, 5}
	x := 5

	idx1, idx2, found := findTwoSumIndices(arr, x)
	if found {
		fmt.Printf("Indices of the elements whose sum is %d: %d, %d\n", x, idx1, idx2)
	} else {
		fmt.Printf("No elements found whose sum is %d\n", x)
	}
}
