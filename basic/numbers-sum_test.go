package basic

import "testing"

func BenchmarkTwoNumbersSumFor(b *testing.B) {
	numbers := []int{5, -2, 10, 8, 6}
	sum := 15
	for i := 0; i < b.N; i++ {
		TwoNumbersSumFor(numbers, sum)
	}
}

func BenchmarkTwoNumbersSumMap(b *testing.B) {
	numbers := []int{5, -2, 10, 8, 6}
	sum := 15
	for i := 0; i < b.N; i++ {
		TwoNumbersSumMap(numbers, sum)
	}
}

func BenchmarkTwoNumbersSumShrinking(b *testing.B) {
	numbers := []int{-2, 2, 3, 7, 9}
	sum := 10
	for i := 0; i < b.N; i++ {
		TwoNumbersSumShrinking(numbers, sum)
	}
}
