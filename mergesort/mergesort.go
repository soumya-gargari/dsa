package main

import "fmt"

func main() {
	data := []int{3, 2, 67, 56, 89, 45, 50}
	result := mergeSort(data)
	fmt.Println(result)
}

func mergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	mid := len(data) / 2
	left := mergeSort(data[:mid])
	right := mergeSort(data[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	result := make([]int, 0)
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	fmt.Println("result is:", result)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
