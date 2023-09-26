package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

// generic print function
func print[T any](v T) { fmt.Println(v) }

// Part 2 sum function refactoring

// numeric interface with a type set
type numeric interface {
	constraints.Integer | constraints.Float
}

// generic sum sums a slice of numbers
func sum[T numeric](numbers ...T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}

func main() {
	// call generic print function
	print("Hello")
	print(42)
	print(true)

	// call generic sum function
	fmt.Println(sum(1, 2, 3))
}
