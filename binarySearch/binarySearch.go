package main

import "fmt"

func binarySearch(data []int, value int) bool {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == value {
			fmt.Println("found value at index:", mid)
			return true
		}
		if value < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	data := []int{2, 3, 5, 4, 89, 45, 67, 93}
	fmt.Println(binarySearch(data, 67))
}
