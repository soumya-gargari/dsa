// our goal is to create two channels, one will take the odd, one will take the even,
// print odd and even numbers concurrently using go routines and channels
package main

import (
	"fmt"
	"sync"
)

func main() {
	oddChan := make(chan int)
	evenChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go sendDataToOddChannel(&wg, oddChan)
	go sendDataToEvenChannel(&wg, evenChan)
	go func() {
		wg.Wait()
		close(oddChan)
		close(evenChan)

	}()
	for i := 0; i < 20; i++ {
		select {
		case oddData, ok := <-oddChan:
			if ok {
				fmt.Println("odd data is:", oddData)
			}
		case evenData, ok := <-evenChan:
			if ok {
				fmt.Println("even data is:", evenData)
			}
		}

	}
}

func sendDataToOddChannel(wg *sync.WaitGroup, oddChan chan<- int) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			oddChan <- i
		}
	}
}

func sendDataToEvenChannel(wg *sync.WaitGroup, evenChan chan<- int) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			evenChan <- i
		}
	}
}
