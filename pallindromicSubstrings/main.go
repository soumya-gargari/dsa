// Given a string s, return the number of palindromic substrings in it.

// A string is a palindrome when it reads the same backward as forward.

// A substring is a contiguous sequence of characters within the string.

// Example 1:

// Input: s = "abc"
// Output: 3
// Explanation: Three palindromic strings: "a", "b", "c".
// Example 2:

// Input: s = "aaa"
// Output: 6
// Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

package main

import "fmt"

func main() {
	data := "aaa"
	fmt.Println("number of palindromic substrings are:", palindromicSubstrings(data))
	fmt.Println("number of palindromic substrings are:", countSubstrings(data))
}

func palindromicSubstrings(s string) int {
	if s == "" {
		return 0
	}
	subStrings := make([]string, 0)
	count := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			subStrings = append(subStrings, s[i:j+1])
		}
	}
	fmt.Println(subStrings)
	for _, val := range subStrings {
		if isPallindrome(val) {
			count++
		}
	}
	return count
}

func isPallindrome(data string) bool {
	temp := []byte(data)
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return data == string(temp)
}

// func countSubstrings(s string) int {
// 	count := 0
// 	for i := range s {
// 		count += isPallindromicSubstring(s, i, i) + isPallindromicSubstring(s, i, i+1)
// 	}
// 	return count
// }

// func isPallindromicSubstring(s string, l, r int) int {
// 	result := 0
// 	for l >= 0 && r < len(s) && s[l] == s[r] {
// 		l--
// 		r++
// 		result++
// 	}
// 	return result
// }

func countSubstrings(data string) int {
	count := 0
	for i := range data {
		count += isPallindromic(data, i, i) + isPallindromic(data, i, i+1)
	}
	return count
}

func isPallindromic(data string, l, r int) int {
	res := 0
	for l >= 0 && r < len(data) && data[l] == data[r] {
		l--
		r++
		res++
	}
	return res
}
