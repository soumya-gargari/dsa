package main

import "fmt"

// You are given two strings s1 and s2 of equal length.
// A string swap is an operation where you choose two indices in a string (not necessarily different) and
// swap the characters at these indices.

// Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings.
// Otherwise, return false.

// Example:
// Input: s1 = "bank", s2 = "kanb"
// Output: true
// Explanation: For example, swap the first character with the last character of s2 to make "bank".

func areAlmostEqual(s1, s2 string) bool {
	first, second, count := -1, -1, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
			if count > 2 {
				return false
			}
		}
		if first == -1 && second == -1 {
			first = i
		} else {
			second = i
		}
	}

	if count == 0 {
		return true
	}
	if first != -1 && second != -1 && s1[first] == s2[second] && s1[second] == s2[first] {
		return true
	}
	return false
}

func main() {
	result := areAlmostEqual("bank", "kanb")
	fmt.Println(result)
}
