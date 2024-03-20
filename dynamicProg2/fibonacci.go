package main

import (
	"fmt"
)

func fibonacci(data int) int {
	fibMap := make(map[int]int)
	result := []int{}
	for i := 0; i <= data; i++ {
		if i < 2 {
			fibMap[i] = i
		} else {
			fibMap[i] = fibMap[i-1] + fibMap[i-2]
		}
	}
	for i := 0; i <= data; i++ {
		result = append(result, fibMap[i])
	}
	fmt.Println("map is:", fibMap)
	fmt.Println("array is:", result)
	return fibMap[data]
}

func main() {
	result := fibonacci(4)
	fmt.Println("result is:", result)
}
