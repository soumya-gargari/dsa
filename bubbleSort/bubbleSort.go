package main

import "fmt"

func main() {
	data := []int{92, 4, 3, 78, 67, 56, 90, 89, 45}
	result := bubbleSort(data)
	fmt.Println("sorted array is:", result)
}

// time complexity is O(n^2)
func bubbleSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}
