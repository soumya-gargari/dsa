package main

import "fmt"

func findMissingNumbers(data []int) []int {
	missingNum := []int{}
	for i := 0; i < len(data); i++ {
		idx := abs(data[i]) - 1
		if data[idx] > 0 {
			data[idx] = data[idx] * -1
		}
	}
	for i := 0; i < len(data); i++ {
		if data[i] > 0 {
			missingNum = append(missingNum, i+1)
		}
	}
	return missingNum
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func main() {
	data := []int{4, 7, 8, 1, 2, 3, 4, 1}
	result := findMissingNumbers(data)
	fmt.Println(result)
}
