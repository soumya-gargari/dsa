package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return

			case stream <- fn():

			}
		}
	}()

	return stream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:

			}
		}

	}()

	return taken
}

func primerFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	isPrime := func(n int) bool {
		for i := n - 1; i > 1; i-- {
			if n%i == 0 {
				return false
			}

		}
		return true
	}
	primeChan := make(chan int)
	go func() {
		defer close(primeChan)
		for {
			select {
			case <-done:
				return
			case randInt := <-randIntStream:
				if isPrime(randInt) {
					fmt.Println("rand int is:", randInt)
					primeChan <- randInt
				}
			}
		}

	}()

	return primeChan
}

func fanIn[T any](done <-chan T, channels []<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)
	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}
	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()
	return fannedInStream
}

func main() {
	//stage 1
	start := time.Now()
	done := make(chan int)
	defer close(done)
	randIntFetcher := func() int { return rand.Intn(50000000) }
	// we will get some value and sent it to a channel
	// randIntStream is channel where we are sending integers
	randIntStream := repeatFunc(done, randIntFetcher)
	// stage 2
	//fanout pattern to create multiple fanout process or goroutines and run it concurrently
	CPUCount := runtime.NumCPU()
	// fmt.Println(CPUCount)
	runtime.GOMAXPROCS(4)
	primeChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeChannels[i] = primerFinder(done, randIntStream)
	}

	// fan in pattern to put all the data into one pipeline and print it
	fannedInStream := fanIn(done, primeChannels)
	for val := range take(done, fannedInStream, 20) {
		fmt.Println("val is:", val)
	}
	fmt.Println(time.Since(start))
}
