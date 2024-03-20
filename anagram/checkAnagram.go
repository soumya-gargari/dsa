package main

import (
	"fmt"
	"sort"
	"strings"
)

// func checkAnagram(data1, data2 string) bool {
// 	if data1 == "" || data2 == "" {
// 		return false
// 	}
// 	if len(data1) == len(data2) {
// 		result1 := strings.Split(data1, "")
// 		result2 := strings.Split(data2, "")
// 		sort.Strings(result1)
// 		sort.Strings(result2)
// 		fmt.Println(result1, result2)
// 		data1 = strings.Join(result1, "")
// 		data2 = strings.Join(result2, "")
// 		fmt.Println(data1, data2)
// 		return data1 == data2
// 	}
// 	return false
// }

func checkAnagram(data1, data2 string) bool {
	a := strings.Split(data1, "")
	b := strings.Split(data2, "")
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	fmt.Println(a, b)
	// sort.Strings(a)
	// sort.Strings(b)
	data1 = strings.Join(a, "")
	data2 = strings.Join(b, "")
	fmt.Println(data1, data2)
	return data1 == data2
}

func main() {
	data1 := "listen"
	data2 := "silent"
	result := checkAnagram(data1, data2)
	fmt.Println(result)
}
