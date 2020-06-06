package main

import (
	"algorithms/basic"
	"fmt"
)

func main() {
	fmt.Printf("[TwoNumbersSumFor] Input: %v want %d -> %v\n", []int{5, -2, 10, 8, 6}, 15, basic.TwoNumbersSumFor([]int{5, -2, 10, 8, 6}, 15))
	fmt.Printf("[TwoNumbersSumMap] Input: %v want %d -> %v\n", []int{5, -2, 10, 8, 6}, 15, basic.TwoNumbersSumFor([]int{5, -2, 10, 8, 6}, 15))
	fmt.Printf("[TwoNumbersSumShrinking] Input: %v want %d -> %v\n", []int{-2, 2, 3, 7, 9}, 10, basic.TwoNumbersSumShrinking([]int{-2, 2, 3, 7, 9}, 10))
	fmt.Printf("[FibonacciRecursive] Input: %d -> %d\n", 6, basic.FibonacciRecursive(6))
	fmt.Printf("[FibonacciDynamic] Input: %d -> %d\n", 6, basic.FibonacciDynamic(6))
	fmt.Printf("[FibonacciSpace] Input: %d -> %d\n", 6, basic.FibonacciSpace(6))
}
