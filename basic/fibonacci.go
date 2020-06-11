package basic

import "fmt"

func FibonacciRecursive(num int) int {
	if num <= 1 {
		return num
	}
	return FibonacciRecursive(num-1) + FibonacciRecursive(num-2)
}

func FibonacciDynamic(num int) int {
	numbers := make([]int, num+2)
	numbers[0] = 0
	numbers[1] = 1
	for i := 2; i <= num; i++ {
		numbers[i] = numbers[i-1] + numbers[i-2]
	}
	return numbers[num]
}

func FibonacciSpace(num int) int {
	a, b, c := 0, 1, 0
	for i := 2; i <= num; i++ {
		c = a + b
		a = b
		b = c
	}
	return b
}

func DemoFibonacci() {
	fmt.Printf("[FibonacciRecursive] Input: %d -> %d\n", 6, FibonacciRecursive(6))
	fmt.Printf("[FibonacciDynamic] Input: %d -> %d\n", 6, FibonacciDynamic(6))
	fmt.Printf("[FibonacciSpace] Input: %d -> %d\n", 6, FibonacciSpace(6))
}
