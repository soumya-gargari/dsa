package main

import "fmt"

func main() {
	data := []int{3, 2, 56, 76, 45, 34, 51}
	fmt.Println("sorted array is:", mergeSort(data))
}

func mergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	mid := len(data) / 2
	leftArray := mergeSort(data[:mid])
	rightArray := mergeSort(data[mid:])
	return merge(leftArray, rightArray)
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
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
