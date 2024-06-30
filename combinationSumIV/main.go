// Given an array of distinct integers nums and a target integer target, return the number of possible combinations that add up to target.

// The test cases are generated so that the answer can fit in a 32-bit integer.

// Example 1:

// Input: nums = [1,2,3], target = 4
// Output: 7
// Explanation:
// The possible combination ways are:
// (1, 1, 1, 1)
// (1, 1, 2)
// (1, 2, 1)
// (1, 3)
// (2, 1, 1)
// (2, 2)
// (3, 1)
// Note that different sequences are counted as different combinations.
// Example 2:

// Input: nums = [9], target = 3
// Output: 0

package main

import "fmt"

func main() {
	input := []int{2, 1, 3}
	target := 4
	fmt.Println("possible combinationn ways to reach the target is:", combinationSum4(input, target))
}

// func combinationSum4(nums []int, target int) int {
// 	var dfs func(target int) int
// 	combinationCount := make(map[int]int, 0)
// 	dfs = func(target int) int {
// 		if target < 0 {
// 			return 0
// 		} else if target == 0 {
// 			return 1
// 		}
// 		if count, exists := combinationCount[target]; exists {
// 			return count
// 		}

// 		count := 0
// 		for _, number := range nums {
// 			count += dfs(target - number)
// 		}
// 		combinationCount[target] = count
// 		fmt.Println(count, combinationCount)
// 		return count
// 	}
// 	return dfs(target)
// }

func combinationSum4(data []int, target int) int {
	var dfs func(target int) int
	dfs = func(target int) int {
		combinationCount := make(map[int]int)
		if target < 0 {
			return 0
		}
		if target == 0 {
			return 1
		}

		if count, ok := combinationCount[target]; ok {
			return count
		}

		count := 0
		for _, num := range data {
			count += dfs(target - num)
		}
		return count
	}
	return dfs(target)
}
