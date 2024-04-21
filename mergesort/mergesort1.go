// package main

// import "fmt"

// func mergeSort(data []int) []int {
// 	if len(data) < 2 {
// 		return data
// 	}
// 	mid := len(data) / 2
// 	left := mergeSort(data[:mid])
// 	right := mergeSort(data[mid:])
// 	return merge(left, right)
// }

// func merge(left, right []int) []int {
// 	result := make([]int, 0, len(left)+len(right))
// 	i, j := 0, 0
// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
// 		}
// 	}
// 	fmt.Println("left is", left[i:])
// 	fmt.Println("right is", right[j:])
// 	result = append(result, left[i:]...)
// 	result = append(result, right[j:]...)
// 	return result
// }

// func main() {
// 	data := []int{4, 2, 3, 78, 56, 45, 90}
// 	result := mergeSort(data)
// 	fmt.Println("sorted array is:", result)
// }
