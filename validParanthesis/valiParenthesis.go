package main

import "fmt"

// func checkIfValid(data string) bool {
// 	stack := make([]rune, 0)
// 	for _, v := range data {
// 		if v == '(' || v == '{' || v == '[' {
// 			stack = append(stack, v)
// 		} else if len(stack) > 0 {
// 			temp := data[len(stack)-1]
// 			if (temp == '{' && v == '}') || (temp == '[' && v == ']') || (temp == '(' && v == ')') {
// 				stack = stack[:len(stack)-1]
// 				fmt.Println(string(stack))
// 			} else {
// 				return false
// 			}
// 		} else {
// 			return false
// 		}
// 	}
// 	return len(stack) == 0
// }

func checkIfValid(data string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(data); i++ {
		if data[i] == '{' || data[i] == '(' || data[i] == '[' {
			stack = append(stack, data[i])
		} else if len(stack) > 0 {
			temp := stack[len(stack)-1]
			if (temp == '{' && data[i] == '}') || (temp == '(' && data[i] == ')') || (temp == '[' && data[i] == ']') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	data := "[{()}"
	fmt.Println(checkIfValid(data))
}
