package basic

import (
	"fmt"
	"testing"
)

var numbers = [4]int{6, 9, 14, 20}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for _, num := range numbers {
		b.Run(fmt.Sprintf("n=%d", num), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibonacciRecursive(num)
			}
		})
	}
}

func BenchmarkFibonacciDynamic(b *testing.B) {
	for _, num := range numbers {
		b.Run(fmt.Sprintf("n=%d", num), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibonacciDynamic(num)
			}
		})
	}
}

func BenchmarkFibonacciSpace(b *testing.B) {
	for _, num := range numbers {
		b.Run(fmt.Sprintf("n=%d", num), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibonacciDynamic(num)
			}
		})
	}
}
