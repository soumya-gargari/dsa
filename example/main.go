package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getResult(c chan string, wg *sync.WaitGroup, ID int) {
	defer wg.Done()
	t := time.Second * time.Duration(rand.Intn(10))
	fmt.Println("Working on", ID, t)
	time.Sleep(t)
	c <- fmt.Sprintf("Result of Req %d is CAT--%d", ID, ID)
	fmt.Println("DONE", ID)
}

func main() {
	workers := 2 // runtime.NumCPU() + 2
	var wg sync.WaitGroup
	results := make(chan string, workers)

	for i := 0; i < workers; i++ {
		fmt.Println("trigger for loop", i)
		wg.Add(1)
		go getResult(results, &wg, i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for val := range results {
		fmt.Println(val)
	}

	fmt.Println("All jobs are completed")
}
