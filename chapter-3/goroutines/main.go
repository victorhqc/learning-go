package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func main() {
	// Add a count of 2, one for each goroutine.
	wg.Add(2)

	fmt.Println("Start goroutines")

	// Launch a goroutine with label "A"
	go printCounts("A")
	// Launch a goroutine with label "B"
	go printCounts("B")

	// Wait for goroutine to finish
	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminating program")
}

func printCounts(label string) {
	// Schedule the call to WaitGroup's Done to tell we are done.
	defer wg.Done()

	// Randomly Wait
	for count := 1; count <= 10; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep))
		fmt.Printf("Count: %d from %s\n", count, label)
	}
}
