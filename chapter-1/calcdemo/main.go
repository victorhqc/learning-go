package main

import (
	"fmt"

	"github.com/victorhqc/learning-go/chapter-1/calc"
)

func main() {
	var x, y int = 10, 5

	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Subtract(x, y))
}
