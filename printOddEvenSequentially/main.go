// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	oddChan := make(chan int)
// 	evenChan := make(chan int)

// 	oddChan1 := make(chan bool, 1)
// 	evenChan1 := make(chan bool, 1)
// 	var wg sync.WaitGroup

// 	defer close(evenChan1)
// 	defer close(oddChan1)
// 	defer close(oddChan)
// 	defer close(evenChan)
// 	wg.Add(4)
// 	go sendDataToOddChan(&wg, oddChan)
// 	go sendDataToEvenChan(&wg, evenChan)
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("odd data is:", <-oddChan)
// 		fmt.Println("even data is:", <-evenChan)
// 	}
// 	go sendEvenData(&wg, evenChan1, oddChan1)
// 	go sendOddData(&wg, evenChan1, oddChan1)
// 	evenChan1 <- true
// 	wg.Wait()
// }

// func sendDataToOddChan(wg *sync.WaitGroup, oddChan chan<- int) {
// 	defer wg.Done()
// 	for i := 1; i <= 20; i++ {
// 		if i%2 == 1 {
// 			oddChan <- i
// 		}
// 	}
// }
// func sendDataToEvenChan(wg *sync.WaitGroup, evenChan chan<- int) {
// 	defer wg.Done()
// 	for i := 1; i <= 20; i++ {
// 		if i%2 == 0 {
// 			evenChan <- i
// 		}
// 	}
// }

// func sendOddData(wg *sync.WaitGroup, evenChan, oddChan chan bool) {
// 	for i := 1; i <= 20; i += 2 {
// 		<-evenChan
// 		fmt.Println("odd val is:", i)
// 		oddChan <- true

// 	}
// 	wg.Done()
// }

// func sendEvenData(wg *sync.WaitGroup, evenChan, oddChan chan bool) {
// 	for i := 2; i <= 20; i += 2 {
// 		<-oddChan
// 		fmt.Println("even val is:", i)
// 		evenChan <- true

// 	}
// 	wg.Done()
// }

// // package main

// // import (
// // 	"fmt"
// // 	"sync"
// // )

// /*
// Problem Statement:
//  Design and implement a Golang program that prints
//  numbers from 1 to 5 in the correct sequence,
//  with even numbers printed by one goroutine and
//  odd numbers printed by a separate goroutine.
// */

// // var wg sync.WaitGroup

// // func main() {
// // 	evenCh, oddCh := make(chan bool, 1), make(chan bool, 1)
// // 	defer close(evenCh)
// // 	defer close(oddCh)

// // 	wg = sync.WaitGroup{}
// // 	wg.Add(2)

// // 	go even(evenCh, oddCh)
// // 	go odd(oddCh, evenCh)

// // 	// initiate the odd function to start first
// // 	evenCh <- true

// // 	wg.Wait()
// // }

// // func even(evenCh, oddCh chan bool) {
// // 	for i := 2; i <= 5; i += 2 {
// // 		<-oddCh
// // 		fmt.Println(i)
// // 		evenCh <- true
// // 	}

// // 	wg.Done()
// // }

// // func odd(oddCh, evenCh chan bool) {
// // 	for i := 1; i <= 5; i += 2 {
// // 		<-evenCh
// // 		fmt.Println(i)
// // 		oddCh <- true
// // 	}

// // 	wg.Done()
// // }

package main

import (
	"fmt"
	"sync"
)

func main() {
	oddChan := make(chan int)
	evenChan := make(chan int)
	oddChan1 := make(chan bool, 1)
	evenChan1 := make(chan bool, 1)
	var wg sync.WaitGroup
	defer close(oddChan)
	defer close(evenChan)
	defer close(oddChan1)
	defer close(evenChan1)
	wg.Add(4)
	go sendDataToOddChannel(&wg, oddChan)
	go sendDataToEvenChannel(&wg, evenChan)

	for i := 0; i < 10; i++ {
		fmt.Println("odd data is:", <-oddChan)
		fmt.Println("even data is:", <-evenChan)
	}
	go sendDataToEvenCh(&wg, oddChan1, evenChan1)
	go sendDataToOddCh(&wg, oddChan1, evenChan1)
	evenChan1 <- true
	wg.Wait()

}

func sendDataToOddChannel(wg *sync.WaitGroup, oddChan chan<- int) {
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			oddChan <- i
		}
	}
	wg.Done()
}

func sendDataToEvenChannel(wg *sync.WaitGroup, evenChan chan<- int) {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			evenChan <- i
		}
	}
	wg.Done()
}

func sendDataToOddCh(wg *sync.WaitGroup, oddChan, evenCh chan bool) {
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			<-evenCh
			fmt.Println("odd val is:", i)
			oddChan <- true
		}
	}
	wg.Done()
}

func sendDataToEvenCh(wg *sync.WaitGroup, oddChan, evenChan chan bool) {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			<-oddChan
			fmt.Println("even val is:", i)
			evenChan <- true
		}
	}
	wg.Done()
}
