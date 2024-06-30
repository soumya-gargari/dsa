// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

// Example 1:

// Input: s = "leetcode"
// Output: 0
// Example 2:

// Input: s = "loveleetcode"
// Output: 2
// Example 3:

// Input: s = "aabb"
// Output: -1

package main

import (
	"fmt"
)

func main() {
	data := "lleetcode"
	fmt.Println("first unique non repeating character index value is:", findNonRepeatingCharacter(data))
}

func findNonRepeatingCharacter(data string) int {
	charSet := make([]int, 26)
	for i := 0; i < len(data); i++ {
		charSet[data[i]-'a'] += 1
	}
	fmt.Println(charSet)
	for i, val := range data {
		if charSet[val-'a'] == 1 {
			return i
		}
	}
	return -1
}

// func findNonRepeatingCharacter(data string) int {
// 	charSet := make([]int, 26)
// 	for i := 0; i < len(data); i++ {
// 		charSet[data[i]-'a'] += 1
// 	}
// 	fmt.Println(charSet)
// 	for i, _ := range charSet {
// 		if charSet[i] == 1 {
// 			return i
// 		}
// 	}
// 	return -1
// }
