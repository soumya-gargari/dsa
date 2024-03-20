package main

import "fmt"

func main() {
	// stockPrices := []int{40, 30, 20, 78, 45, 89, 34}
	// min := stockPrices[0]
	// profit := 0
	// for i := 1; i < len(stockPrices); i++ {
	// 	if stockPrices[i] < min {
	// 		min = stockPrices[i]
	// 	} else {
	// 		temp := stockPrices[i] - min
	// 		if temp > profit {
	// 			profit = temp
	// 		}

	// 	}
	// }

	stockPrices := []int{20, 10, 32, 52, 8, 10, 34, 51}
	profit := 0
	min := stockPrices[0]
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
