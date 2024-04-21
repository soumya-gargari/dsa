package main

import "fmt"

// get max sum within a window provided as i/p
// sliding window alogo where we are shifting the window one by one based on the window size provided as input
// and calculate max sum of that window subarray and print it
// func slidingWindow(data []int, k int) int {
// 	var maxSum, startIndex, endIndex int
// 	sum := 0
// 	for i := 0; i < k; i++ {
// 		sum += data[i]
// 	}
// 	maxSum = sum
// 	for i := k; i < len(data); i++ {
// 		sum = sum + data[i] - data[i-k]
// 		if sum > maxSum {
// 			maxSum = sum
// 			startIndex = i - k + 1
// 			endIndex = i
// 		}
// 	}
// 	fmt.Println(startIndex, endIndex)
// 	return maxSum
// }

func slidingWindow(data []int, k int) int {
	sum, max := 0, 0
	for i := 0; i < k; i++ {
		sum += data[i]
	}
	max = sum
	for i := k; i < len(data); i++ {
		sum = sum + data[i-k]
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	data := []int{2, 3, -6, 4, 10, 1, -8, 3}
	result := slidingWindow(data, 4)
	fmt.Println(result)
}
