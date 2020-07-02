package basic

import (
	"fmt"
	"testing"
)

var fibNumbers = [4]int{6, 9, 14, 20}

type fibonacciTest struct {
	num  int
	want int
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for _, num := range fibNumbers {
		b.Run(fmt.Sprintf("n=%d", num), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibonacciRecursive(num)
			}
		})
	}
}

func BenchmarkFibonacciDynamic(b *testing.B) {
	for _, num := range fibNumbers {
		b.Run(fmt.Sprintf("n=%d", num), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibonacciDynamic(num)
			}
		})
	}
}

func BenchmarkFibonacciSpace(b *testing.B) {
	for _, num := range fibNumbers {
		b.Run(fmt.Sprintf("n=%d", num), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibonacciDynamic(num)
			}
		})
	}
}

func testModels() []fibonacciTest {
	return []fibonacciTest{
		{num: 6, want: 8},
		{num: 7, want: 13},
		{num: 8, want: 21},
		{num: 9, want: 34},
		{num: 10, want: 55},
		{num: 11, want: 89},
	}
}

func TestFibonacciRecursive(t *testing.T) {
	tests := testModels()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.num), func(t *testing.T) {
			if got := FibonacciRecursive(tt.num); got != tt.want {
				t.Errorf("FibonacciRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciFibonacciDynamic(t *testing.T) {
	tests := testModels()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.num), func(t *testing.T) {
			if got := FibonacciDynamic(tt.num); got != tt.want {
				t.Errorf("FibonacciDynamic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciSpace(t *testing.T) {
	tests := testModels()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.num), func(t *testing.T) {
			if got := FibonacciSpace(tt.num); got != tt.want {
				t.Errorf("FibonacciSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
