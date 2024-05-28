package main

import (
	"fmt"
)

func main() {
	w1 := "abcdafsd"
	w2 := "acbcfsda"
	result := findLongedstCommonSubsequence(w1, w2)
	fmt.Println(result)
}

func findLongedstCommonSubsequence(w1, w2 string) int {
	l1 := len(w1)
	l2 := len(w2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// time complexity is O(m*n) where m and n came for making 2D array, space complexity is also same..
// func findLongedstCommonSubsequence(w1, w2 string) int {
// 	l1 := len(w1)
// 	l2 := len(w2)
// 	dp := make([][]int, l1+1)
// 	for i := 0; i <= l1; i++ {
// 		dp[i] = make([]int, l2+1)
// 	}
// 	for i := 1; i <= l1; i++ {
// 		for j := 1; j <= l2; j++ {
// 			if w1[i-1] == w2[j-1] {
// 				// if same val, then choose diagonal opposite value and add 1 to it
// 				dp[i][j] = 1 + dp[i-1][j-1]
// 			} else {
// 				// compare top and left value and store it
// 				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
// 			}
// 		}
// 	}
// 	return dp[l1][l2]
// }
