// package main

// import "fmt"

// func main() {
// 	w1, w2 := "abcfghjkuil", "jkuil"

// 	result := longestCommonString(w1, w2)
// 	fmt.Println("longest common string is:", result)
// 	result = longestCommonString("dfghjabcba", reverse("abcba"))
// 	fmt.Println("longest common pallindromic string is:", result)
// 	a := []int{4, 5, 6, 3, 4, 6, 90}
// 	b := []int{3, 4, 6, 90, 89, 78}
// 	result1 := longestCommonSubArray(a, b)
// 	fmt.Println("longest common subarray is:", result1)

// }

// func longestCommonString(w1, w2 string) string {
// 	l1 := len(w1)
// 	l2 := len(w2)
// 	lcString := ""
// 	var row, col int
// 	dp := make([][]int, l1+1)
// 	for i := 0; i <= l1; i++ {
// 		dp[i] = make([]int, l2+1)
// 	}
// 	for i := 1; i <= l1; i++ {
// 		for j := 1; j <= l2; j++ {
// 			if w1[i-1] == w2[j-1] {
// 				dp[i][j] = 1 + dp[i-1][j-1]
// 				row = i
// 				col = j
// 			}
// 		}
// 	}
// 	for dp[row][col] != 0 {
// 		lcString += string(w1[row-1])
// 		row--
// 		col--
// 	}
// 	fmt.Println(lcString)
// 	op := reverse(lcString)
// 	return op
// }

// func reverse(data string) string {
// 	result := []byte(data)
// 	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
// 		result[i], result[j] = result[j], result[i]
// 	}
// 	return string(result)
// }

// func reverseIntArray(data []int) []int {
// 	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
// 		data[i], data[j] = data[j], data[i]
// 	}
// 	return data
// }

// func longestCommonSubArray(w1, w2 []int) []int {
// 	l1 := len(w1)
// 	l2 := len(w2)
// 	lcSubArray := []int{}
// 	var row, col int
// 	dp := make([][]int, l1+1)
// 	for i := 0; i <= l1; i++ {
// 		dp[i] = make([]int, l2+1)
// 	}
// 	for i := 1; i <= l1; i++ {
// 		for j := 1; j <= l2; j++ {
// 			if w1[i-1] == w2[j-1] {
// 				dp[i][j] = 1 + dp[i-1][j-1]
// 				row = i
// 				col = j
// 			}
// 		}
// 	}
// 	for dp[row][col] != 0 {
// 		lcSubArray = append(lcSubArray, w1[row-1])
// 		row--
// 		col--
// 	}
// 	fmt.Println(lcSubArray)
// 	op := reverseIntArray(lcSubArray)
// 	return op
// }

package main

import "fmt"

func findLongestCommonString(w1, w2 string) string {
	l1 := len(w1)
	l2 := len(w2)
	lString := ""
	var row, col int
	dp := make([][]int, l1+1)
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
	fmt.Println("string is:", lString)
	result := reverse(lString)
	return result
}

func reverse(data string) string {
	result := []byte(data)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func main() {
	s1 := "abghefnlj"
	s2 := "kevkrkghefnljnk"
	result := findLongestCommonString(s1, s2)
	fmt.Println("longest common string is:", result)
}
