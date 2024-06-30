// You are given an integer array coins representing coins of different denominations and an integer amount representing
// a total amount of money.

// Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by
// any combination of the coins, return -1.

// You may assume that you have an infinite number of each kind of coin.

// Example 1:

// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
// Example 2:

// Input: coins = [2], amount = 3
// Output: -1
// Example 3:

// Input: coins = [1], amount = 0
// Output: 0

package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{1, 2, 5}
	target := 10
	fmt.Println("minimum number of coins to find the target is:", findMinCoins(coins, target))
}

// Time complexity : O(m *n) where m is number of coins and n is amount to make
// space complexity : O(n)
// func findMinCoins(coins []int, target int) int {
// 	if target < 1 {
// 		return 0
// 	}
// 	dp := make([]int, target+1)
// 	for i := 1; i <= target; i++ {
// 		dp[i] = math.MaxInt32
// 	}
// 	for _, coin := range coins {
// 		for a := 1; a <= target; a++ {
// 			if a >= coin {
// 				dp[a] = Min(dp[a], 1+dp[a-coin])
// 			}
// 		}
// 	}
// 	return dp[target]
// }

// func Min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func findMinCoins(coins []int, target int) int {
	if target < 1 {
		return 0
	}
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = math.MaxInt32
	}

	for _, coin := range coins {
		for a := 1; a <= target; a++ {
			if a >= coin {
				// Min needs to to between dp[a] and 1+dp[a-coin] because suppose 5 u need to make
				// in array, 5 coin us already there, then you dont need 1+ dp[a-coin], u need dp[a], 
				// because in that case dp[a] is minimum, thats why
				dp[a] = Min(dp[a], 1+dp[a-coin])
			}
		}
	}
	return dp[target]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
