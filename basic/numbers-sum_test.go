package basic

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func generateRandomSlice(num int, limit int) []int {
	numbers := make([]int, 0, num)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		numbers = append(numbers, rand.Intn(limit))
	}
	return numbers
}

func BenchmarkTwoNumbersSumFor(b *testing.B) {
	numbers := generateRandomSlice(12345, 15)
	sum := 15
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoNumbersSumFor(numbers, sum)
	}
}

func BenchmarkTwoNumbersSumMap(b *testing.B) {
	numbers := generateRandomSlice(12345, 15)
	sum := 15
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoNumbersSumMap(numbers, sum)
	}
}

func BenchmarkTwoNumbersSumShrinking(b *testing.B) {
	numbers := generateRandomSlice(12345, 20)
	sort.Ints(numbers)
	sum := 20
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoNumbersSumShrinking(numbers, sum)
	}
}
