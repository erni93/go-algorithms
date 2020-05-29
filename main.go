package main

import (
	"algorithms/basic"
	"fmt"
)

func main() {
	fmt.Printf("[TwoNumbersSumFor] Input: %v -> %v\n", []int{5, -2, 10, 8, 6}, basic.TwoNumbersSumFor([]int{5, -2, 10, 8, 6}, 15))
	fmt.Printf("[TwoNumbersSumMap] Input: %v -> %v\n", []int{5, -2, 10, 8, 6}, basic.TwoNumbersSumFor([]int{5, -2, 10, 8, 6}, 15))
	fmt.Printf("[TwoNumbersSumMap] Input: %v -> %v\n", []int{-2, 2, 3, 7, 9}, basic.TwoNumbersSumShrinking([]int{-2, 2, 3, 7, 9}, 10))
}
