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
