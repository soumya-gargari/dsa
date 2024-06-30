// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a
// space-separated sequence of one or more dictionary words.

// Note that the same word in the dictionary may be reused multiple times in the segmentation.

// Example 1:

// Input: s = "leetcode", wordDict = ["leet","code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".
// Example 2:

// Input: s = "applepenapple", wordDict = ["apple","pen"]
// Output: true
// Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
// Note that you are allowed to reuse a dictionary word.
// Example 3:

// Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// Output: false
package main

import (
	"fmt"
	//"sort"
)

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	fmt.Println("given string can be segmented into one or more dictionary words:", wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	max_len := 0
	for _, word := range wordDict {
		if len(word) > max_len {
			max_len = len(word)
		}
	}
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= max(i-max_len-1, 0); j-- {
			fmt.Println(i, j)
			if dp[j] && contains(wordDict, s[j:i]) {
				fmt.Println(s[j:i])

				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func contains(wordDict []string, word string) bool {
	for _, val := range wordDict {
		if val == word {
			return true
		}
	}
	return false
}
