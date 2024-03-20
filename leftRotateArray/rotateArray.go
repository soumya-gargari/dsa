package main

import "fmt"

func rotateArray1(data []int, k int) {
	k %= len(data)
	temp := make([]int, 0)
	for i := 0; i < k; i++ {
		temp = append(temp, data[i])
	}
	data = data[k:]
	fmt.Println(data)
	data = append(data, temp...)
	fmt.Println("rotated array is:", data)
}

func main() {
	data := []int{2, 3, 4, 5, 6, 7}
	k := 0
	fmt.Printf("Enter the number how many times you want to rotate the array:")
	fmt.Scanf("%d", &k)
	rotateArray1(data, k)
}
