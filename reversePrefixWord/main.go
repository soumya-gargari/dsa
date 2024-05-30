// Given a 0-indexed string word and a character ch, reverse the segment of word that starts at index 0 and ends at the index of the first occurrence of ch (inclusive). If the character ch does not exist in word, do nothing.

// For example, if word = "abcdefd" and ch = "d", then you should reverse the segment that starts at 0 and ends at 3 (inclusive). The resulting string will be "dcbaefd".
// Return the resulting string.

// Example 1:

// Input: word = "abcdefd", ch = "d"
// Output: "dcbaefd"
// Explanation: The first occurrence of "d" is at index 3.
// Reverse the part of word from 0 to 3 (inclusive), the resulting string is "dcbaefd".
// Example 2:

// Input: word = "xyxzxe", ch = "z"
// Output: "zxyxxe"
// Explanation: The first and only occurrence of "z" is at index 3.
// Reverse the part of word from 0 to 3 (inclusive), the resulting string is "zxyxxe".
// Example 3:

// Input: word = "abcd", ch = "z"
// Output: "abcd"
// Explanation: "z" does not exist in word.
// You should not do any reverse operation, the resulting string is "abcd".

package main

import (
	"fmt"
)

func main() {
	data := "abcdefd"
	character := 'd'
	fmt.Println("reversed prefix word is:", reversePrefixWord(data, byte(character)))
}

func reversePrefixWord(word string, ch byte) string {
	stack := make([]byte, 0)
	var result string
	var match bool
	breakpoint := 0
	if word[0] == ch {
		return word
	}
	for i := 0; i < len(word); i++ {
		if word[i] != ch {
			stack = append(stack, word[i])
		} else if word[i] == ch {
			match = true
			breakpoint = i
			break
		}
	}
	if match {
		partition := word[breakpoint+1:]
		stack = append(stack, ch)
		result = reverse(stack)
		return result + partition
	}
	return word
}

func reverse(data []byte) string {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return string(data)
}
