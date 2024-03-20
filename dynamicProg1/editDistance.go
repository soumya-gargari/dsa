// In this problems, we will decide how many operations are needed to convert word1 to word2
// for ex: change "horse" to "ros"
// o/p will be 3 as 3 operations are needed to change horse to ros

// Approach is : make 2D array for storing each letter exchange operations
// initialize with extra length of 1 because of the empty strings
// check top left, left and top 2D array data which is minimum, and add 1 to it and store it in 2D array.
package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func editDistance(w1, w2 string) int {
	l1 := len(w1)
	l2 := len(w2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	// made extra 1 of length in 2D array because of empty strings

	// below, storing all the data of how ,many operations are needed to convert w1 to w2
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// topleft = dp[i-1][j-1]
				// left = dp[i][j-1]
				// top = dp[i-1][j]
				// compare these 3 values, choose min value and add 1 to it
				dp[i][j] = 1 + min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
			}
		}
	}
	return dp[l1][l2]
}

func main() {
	w1 := "abcdef"
	w2 := "azcdegvh"
	data := editDistance(w1, w2)
	fmt.Println("operations needed to convert w1 to w2 is:", data)
}
