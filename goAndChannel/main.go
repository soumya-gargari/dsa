// package main

// /*
// 	Goal: Printing Odd numbers and Even Numbers sequentially
// 		1. Usage of golang channels to print odd and even numbers
// 		2. Understand the Logic for Odd and even Number find out
// */

// import "fmt"

// // Print - This helps in selecting the odd and even numbers sequentially
// func Print(odd, even <-chan int) {
// 	for {
// 		select {
// 		case item := <-odd:
// 			fmt.Println(item)
// 		case item := <-even:
// 			fmt.Println(item)
// 		}
// 	}
// }

// // PrintOddNum - This helps in selecting the odd numbers sequentially
// func PrintOddNum(odd <-chan int) {
// 	for {
// 		select {
// 		case item := <-odd:
// 			fmt.Println("odd item is:", item)
// 		}
// 	}
// }

// // PrintEvenNum - This helps in selecting even numbers sequentially
// func PrintEvenNum(even <-chan int) {
// 	for {
// 		select {
// 		case item := <-even:
// 			fmt.Println("even item is:", item)
// 		}
// 	}
// }

// func main() {
// 	oddChan := make(chan int)
// 	evenChan := make(chan int)
// 	go Print(oddChan, evenChan)
// 	for i := range [20]int{} {
// 		if i%2 == 0 {
// 			oddChan <- i
// 		} else {
// 			evenChan <- i
// 		}
// 	}
// 	go PrintOddNum(oddChan)
// 	go PrintEvenNum(evenChan)
// 	for i := range [20]int{} {
// 		if i%2 == 0 {
// 			oddChan <- i
// 		} else {
// 			evenChan <- i
// 		}
// 	}
// }

package main

import (
	"fmt"
	"sync"
)

// PrintOddNum - This helps in selecting the odd numbers sequentially
func PrintOddNum(odd chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			odd <- i
		}
	}
}

// PrintEvenNum - This helps in selecting even numbers sequentially
func PrintEvenNum(even chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			even <- i
		}
	}
}

func main() {
	oddChan := make(chan int)
	evenChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go PrintOddNum(oddChan, &wg)
	go PrintEvenNum(evenChan, &wg)

	for i := 0; i < 10; i++ {
		fmt.Println("odd number is:", <-oddChan)
		fmt.Println("even number is:", <-evenChan)
	}
	go func() {
		defer wg.Wait()
		close(oddChan)
		close(evenChan)
	}()
}
