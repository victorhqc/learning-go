package main

import (
	"fmt"

	"github.com/victorhqc/learning-go/chapter-2/strcon"
)

func main() {
	s := strcon.SwapCase("Gopher")
	fmt.Println("Converted string is:", s)
}
