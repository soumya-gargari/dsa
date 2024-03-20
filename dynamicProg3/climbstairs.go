package main

import "fmt"

func climbstairs(numOfStairs int) int {
	result := make(map[int]int)
	result[1] = 1
	result[2] = 2

	for i := 3; i <= numOfStairs; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[numOfStairs]
}

func main() {
	numOfStairs := 5
	waysToTakeDest := climbstairs(numOfStairs)
	fmt.Println("to reach those number of stairs, we can perform operations of: ", waysToTakeDest)
}
