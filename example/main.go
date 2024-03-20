// code for checking pallindrome
package main

import (
	"fmt"
)

func main() {
	input := "madabngabbcddg"
	outPut := printLongestSubString(input)
	fmt.Println("longest substring without repeating characters are:", outPut)

}

// abbcdb
func printLongestSubString(data string) string {
	charSet := make(map[byte]int, 0)
	start := 0
	maxLength := 0
	longestSubstart := 0
	for end := 0; end < len(data); end++ {
		if _, ok := charSet[data[end]]; ok {
			start = maxVal(start, charSet[data[end]]+1)
		}
		charSet[data[end]] = int(data[end])
		if end-start+1 > maxLength {
			maxLength = end - start + 1
			longestSubstart = start
		}
	}
	fmt.Println(longestSubstart, maxLength)
	output := data[longestSubstart : longestSubstart+maxLength]
	return output
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
