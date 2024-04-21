// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.

// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.
// Example 2:
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.
package main

import "fmt"

func findThreesumTriplets(data []int) [][]int {
	results := make([][]int, 0)
	if data == nil || len(data) < 3 {
		return results
	}
	for i := 0; i < len(data)-2; i++ {
		// fix 1st element, take left pointer(starting from 2nd index) and right pointer(which is starting from last index)
		// and do the addition, if matches then return the values, if <0, then increase left pointer, else decrese right pointer
		left := i + 1
		right := len(data) - 2
		sum := data[i] + data[left] + data[right]
		if sum == 0 {
			result := make([]int, 0)
			result = append(result, data[i], data[left], data[right])
			results = append(results, result)
			left++
			right--
		} else if sum < 0 {
			left++
		} else {
			right--
		}

	}
	return results

}

func main() {
	data := []int{-1, 0, 1, 2, -1, 4}
	result := findThreesumTriplets(data)
	fmt.Println("triplets are:", result)
}
