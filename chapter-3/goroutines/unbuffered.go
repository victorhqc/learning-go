package main

import (
	"fmt"
	"sync"
)

// wg is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func main() {
	count := make(chan int)

	// Add a count of 2, one for each goroutine.
	wg.Add(2)
	fmt.Println("Start goroutines.")

	// Launch a goroutine with label "A"
	go printUnbufferedCounts("A", count)
	// Launch a goroutine with label "B"
	go printUnbufferedCounts("B", count)

	fmt.Println("Channel begin")
	count <- 1

	// Wait for goroutines to finish.
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func printUnbufferedCounts(label string, count chan int) {
	// Schedule the call to WaitGroup's done to tell we are.
	defer wg.Done()

	for {
		// Receives a message from channel
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed")
			return
		}

		fmt.Printf("Count: %d received from %s\n", val, label)
		if val == 10 {
			fmt.Printf("Channel closed from %s\n", label)
			// Close the channel
			close(count)
			return
		}
		val++
		// Send back to the other goroutine.
		count <- val
	}
}
