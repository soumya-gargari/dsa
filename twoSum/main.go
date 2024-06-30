// Given an input array and a integer n, we have two tell in how many ways we can get number n by summing the numbers of given array,
// we can use a number only once.
// example: array[5] = {1,3,7,9,10,11} and n = 12 then
// possible cases are 1+11, 9+3. so ans = 2.

package main

import "fmt"

func main() {
	array := []int{1, 3, 7, 9, 10, 11, 5}
	target := 12
	result := countSubsetsDP(array, target)
	fmt.Printf("Number of ways to sum to %d: %d\n", target, result)
}

func countSubsetsDP(data []int, target int) int {
	if len(data) == 0 || len(data) < 2 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range data {
		// traverse in reversal way to prevent using same elements
		for j := target; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
		}
	}
	fmt.Println(dp)
	return dp[target]
}
