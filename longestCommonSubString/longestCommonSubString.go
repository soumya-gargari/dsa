package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	start := 0
	maxLength := 0

	for end := 0; end < len(s); end++ {
		if _, ok := charIndex[s[end]]; ok {
			start = max(start, charIndex[s[end]]+1)
		}
		charIndex[s[end]] = end
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func longestSubstring(data string) string {
// 	start := 0
// 	maxLength := 0
// 	longestStart := 0
// 	charSet := make(map[byte]int)
// 	for end := 0; end < len(data); end++ {
// 		if _, ok := charSet[data[end]]; ok {
// 			start = max(start, charSet[data[end]]+1)
// 		}
// 		charSet[data[end]] = end
// 		if end-start+1 > maxLength {
// 			maxLength = end - start + 1
// 			longestStart = start
// 			fmt.Println("first index of longest substring is:", longestStart)
// 		}
// 	}
// 	l := data[longestStart : longestStart+maxLength]
// 	fmt.Printf("longest substring is:%v and its length is %d\n", l, maxLength)
// 	return l
// }

func longestSubstring(data string) string {
	charSet := make(map[byte]int, 0)
	start := 0
	maxLength := 0
	longestSubStart := 0
	for end := 0; end < len(data); end++ {
		if _, ok := charSet[data[end]]; ok {
			start = max(start, charSet[data[end]]+1)
		}
		charSet[data[end]] = end
		if end-start+1 > maxLength {
			maxLength = end - start + 1
			longestSubStart = start
		}
	}
	result := data[longestSubStart : longestSubStart+maxLength]
	return result
}

func main() {
	input := "abcabcbfasd"
	//length := lengthOfLongestSubstring(input)
	longestSubstring := longestSubstring(input)

	//fmt.Printf("Length of the longest substring without repeating characters: %d\n", length)
	fmt.Printf("Longest substring without repeating characters: %s\n", longestSubstring)
}
