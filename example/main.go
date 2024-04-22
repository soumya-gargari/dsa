package main

import (
	"fmt"
)

func main() {
	w1, w2 := "abcfghjk", "fghjkl"
	result := findLongestCommonString(w1, w2)
	fmt.Println("longest common string is:", result)
	w3, w4 := "abcfghjkyu", "fghjklyui"
	resultLength := findLongestCommonSequence(w3, w4)
	fmt.Println("longest common sequence is:", resultLength)
	data := "hbfwbbjabcfghjkll"
	fmt.Println("longest common sub string is:", findLongestCommonSubString(data))
	datas := []int{3, 5, 45, 67, 89, 90}
	fmt.Println("is data found:", binarySearch(datas, 6))

	// sliding window algo
	arrays := []int{-2, -1, 9, 5, -3, 9, 1, 8}
	fmt.Println("max sum within the window is:", slidingWindow(arrays, 4))
}

func slidingWindow(data []int, k int) int {
	maxSum := 0
	sum := 0
	for i := 0; i < k; i++ {
		sum += data[i]
	}
	maxSum = sum
	for i := k; i < len(data); i++ {
		sum = sum - data[i-k] + data[i]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func binarySearch(data []int, val int) bool {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == val {
			fmt.Println("found at index:", mid)
			return true
		}
		if val < data[mid] {
			high--
		} else {
			low++
		}
	}
	return false
}

func findLongestCommonSubString(data string) string {
	maxLength := 0
	start := 0
	lSubStart := 0
	charSet := make(map[byte]int, 0)
	for end := 0; end < len(data); end++ {
		if _, ok := charSet[data[end]]; ok {
			start = max(start, charSet[data[end]]+1)
		}
		charSet[data[end]] = end
		if end-start+1 > maxLength {
			maxLength = end - start + 1
			lSubStart = start
		}
	}
	return data[lSubStart:(lSubStart + maxLength)]
}

func findLongestCommonString(w1, w2 string) string {
	l1 := len(w1)
	l2 := len(w2)
	lString := ""
	dp := make([][]int, l1+1)
	var row, col int
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				row = i
				col = j
			}
		}
	}
	for dp[row][col] != 0 {
		lString += string(w1[row-1])
		row--
		col--
	}
	return reverse(lString)
}

func findLongestCommonSequence(w1, w2 string) int {
	l1 := len(w1)
	l2 := len(w2)
	dp := make([][]int, l1+1)
	//	var row, col int
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = maximum(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(data string) string {
	runes := ([]byte(data))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
