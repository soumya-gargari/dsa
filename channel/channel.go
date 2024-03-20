package main

import (
	"fmt"
	"sync"
)

func write(ch chan int) {
	//wg.Add(1)
	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	//defer wg.Done()
	close(ch)
}

func send(wg *sync.WaitGroup, ch chan<- int) {
	wg.Add(1)
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("value sent", i)
	}
	close(ch)
	defer wg.Done()
}

func receive(ch <-chan int) {

	for v := range ch {
		fmt.Println("value received:", v)
	}

}

func main() {

	// creates capacity of 2
	var wg sync.WaitGroup
	ch := make(chan int)
	go write(ch)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
	}
	ch1 := make(chan int)
	// data := []int{2, 3, 4, 5, 6}
	go send(&wg, ch1)
	receive(ch1)
	wg.Wait()
}
