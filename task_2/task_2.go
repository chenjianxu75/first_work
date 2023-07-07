package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	// Create buffered channels with a capacity of 1 to control the order of printing
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)
	ch4 := make(chan struct{}, 1)

	// Create a mutex and a counter to track the number of prints
	var mutex sync.Mutex
	counter := 0

	// Start the goroutines with proper synchronization
	go printName("Zhang San", ch1, ch2, &wg, &mutex, &counter)
	go printName("Li Si", ch2, ch3, &wg, &mutex, &counter)
	go printName("Wang Wu", ch3, ch4, &wg, &mutex, &counter)
	go printName("Zhao Liu", ch4, ch1, &wg, &mutex, &counter)

	// Send an initial value to the first channel to start the printing process
	ch1 <- struct{}{}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the channels
	close(ch1)
	close(ch2)
	close(ch3)
	close(ch4)
}

func printName(name string, ch <-chan struct{}, nextCh chan<- struct{}, wg *sync.WaitGroup, mutex *sync.Mutex, counter *int) {
	defer wg.Done()

	for range ch {
		// Acquire the mutex to update the counter
		mutex.Lock()
		*counter++
		printCount := *counter
		mutex.Unlock()

		// Print the name and the current count
		fmt.Printf("%s (%d)\n", name, printCount)

		// Exit the loop if the maximum count is reached
		if printCount >= 400 {
			break
		}

		// Send a value to the next channel to allow the next goroutine to proceed
		nextCh <- struct{}{}
	}
}
