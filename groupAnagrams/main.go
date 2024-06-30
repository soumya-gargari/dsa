// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:

// Input: strs = [""]
// Output: [[""]]
// Example 3:

// Input: strs = ["a"]
// Output: [["a"]]

package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := findGroupAnagrams(data)
	fmt.Println("group anagrams are :", result)
}

// in this method, time complexity is O(nlogk) where k is max length of strings and n is number of strings
// func findGroupAnagrams(data []string) [][]string {
// 	anagramMap := make(map[string][]string, 0)
// 	results := make([][]string, 0)
// 	if data == nil {
// 		return results
// 	}
// 	for i := 0; i < len(data); i++ {
// 		temp := []byte(data[i])
// 		sort.Slice(temp, func(i, j int) bool {
// 			return temp[i] < temp[j]
// 		})
// 		anagramMap[string(temp)] = append(anagramMap[string(temp)], data[i])
// 	}
// 	fmt.Println(anagramMap)
// 	for _, val := range anagramMap {
// 		result := make([]string, 0)
// 		result = append(result, val...)
// 		results = append(results, result)
// 	}
// 	fmt.Println(results)
// 	return results
// }

func findGroupAnagrams(data []string) [][]string {
	results := make([][]string, 0)
	if len(data) == 0 || data[0] == "" {
		return results
	}
	anagramMap := make(map[string][]string, 0)

	for i := 0; i < len(data); i++ {
		temp := []byte(data[i])
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		anagramMap[string(temp)] = append(anagramMap[string(temp)], data[i])
	}
	fmt.Println(anagramMap)
	for _, val := range anagramMap {
		result := make([]string, 0)
		result = append(result, val...)
		results = append(results, result)
	}
	return results
}
