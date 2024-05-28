package main

import (
	"fmt"
	"sort"
	"strings"
)

func checkAnagram(data1, data2 string) bool {
	w1 := strings.Split(data1, "")
	sort.Slice(w1, func(i, j int) bool {
		return w1[i] < w1[j]
	})
	w2 := strings.Split(data2, "")
	sort.Slice(w2, func(i, j int) bool {
		return w2[i] < w2[j]
	})
	fmt.Println(w1, w2)
	data1 = strings.Join(w1, "")
	data2 = strings.Join(w2, "")
	return data1 == data2
}

func main() {
	data1 := "listen"
	data2 := "silent"
	result := checkAnagram(data1, data2)
	fmt.Println(result)
}
