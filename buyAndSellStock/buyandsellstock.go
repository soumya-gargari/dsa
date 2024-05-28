package main

import "fmt"

func main() {
	stockPrices := []int{8, 9, 10, 7, 15, 26, 27}
	min := stockPrices[0]
	profit := 0
	for i := 1; i < len(stockPrices); i++ {
		if min > stockPrices[i] {
			min = stockPrices[i]
		} else {
			temp := stockPrices[i] - min
			if temp > profit {
				profit = temp
			}
		}
	}

	fmt.Println("max profit is:", profit)

}
