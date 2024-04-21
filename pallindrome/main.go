package main

import "fmt"

func main() {
	data := 121
	fmt.Println(isPallindrome(data))
}

func isPallindrome(data int) bool {
	reversedData := reverse(data)
	return reversedData == data
}

func reverse(data int) int {
	reversedData := 0
	for data > 0 {
		remainder := data % 10
		reversedData = reversedData*10 + remainder
		data = data / 10
	}
	return reversedData
}
