package main

import (
	"fmt"
)

func main() {
	data := 9
	fmt.Println("how many times consequtive numbers can be added and return the input:", consecutiveNumbersSum(data))
}

// x+(x+1)+(x+2)+...+(x+k−1)=N
// x+0+x+1+x+2+...+x+k−1=N
// x∗k+(0+1+2+3+...+k−1)=N
// (0+1+2+3+...+k−1)(0 + 1 + 2 + 3 + ... + k - 1)(0+1+2+3+...+k−1) is the sum of numbers from 0 to k - 1, which we can write as k∗(k−1)/2k * (k - 1) / 2k∗(k−1)/2
// And, then
// x∗k+k∗(k−1)/2= N
// Hence, represent the x
// Divine By K
// x+(k−1)/2=N/k
// x + (k - 1)/2 = N/k
// x+(k−1)/2=N/k
// x= (n− k(k−1)/2)/k​

// func consecutiveNumbersSum(n int) int {

// 	res := 0
// 	for k := 1; k*(k-1)/2 < n; k++ {
// 		if (n-k*(k-1)/2)%k == 0 {
// 			res++
// 		}
// 	}
// 	return res
// }

func consecutiveNumbersSum(n int) int {
	count := 0
	if n < 1 {
		return 0
	}
	for k := 1; k < k*(k-1)/2; k++ {
		if (n-(k*(k-1)/2))%k == 0 {
			count++
		}
	}
	return count
}
