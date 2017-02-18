package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50}

	for key, value := range x {
		fmt.Printf("Index: %d Value: %d\n", key, value)
	}
}
