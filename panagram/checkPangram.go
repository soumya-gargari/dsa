package main

import (
	"fmt"
)

// check if in a given string. all the alphabets from "a" to "z" are present or not
// if yes, return true else return false

func checkPangram(data string) bool {
	result := make(map[int]string)
	// making a result map where in each index , all the alphabets will be present sequentially
	for i := 'a'; i <= 'z'; i++ {
		result[int(i-'a')] = string(i)
	}
	fmt.Println(result)
	for i := 0; i < len(data); i++ {
		delete(result, int(data[i]-'a'))
		if len(result) == 0 {
			return true
		}
	}
	fmt.Println("map is;", result)
	return false
}

func main() {
	data := "thequickbrownfoxjumpsoverthelazydg"
	fmt.Println("result is:", checkPangram(data))
}
