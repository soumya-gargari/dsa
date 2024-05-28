package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubstring(data string) string {
	charSet := make(map[byte]int)
	start := 0
	longestSubStart := 0
	maxLen := 0
	for end := 0; end < len(data); end++ {
		if _, ok := charSet[data[end]]; ok {
			start = max(start, charSet[data[end]]+1)
		}
		charSet[data[end]] = end
		if end-start+1 > maxLen {
			maxLen = end - start + 1
			longestSubStart = start
		}
	}
	return data[longestSubStart : longestSubStart+maxLen]
}

func main() {
	input := "abcabcbfasdaaa"
	//length := lengthOfLongestSubstring(input)
	longestSubstring := longestSubstring(input)

	//fmt.Printf("Length of the longest substring without repeating characters: %d\n", length)
	fmt.Printf("Longest substring without repeating characters: %s\n", longestSubstring)
}
