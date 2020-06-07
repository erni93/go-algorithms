package basic

import (
	"testing"
)

func generateNumbers(start int, quantity int) []int {
	numbers := make([]int, quantity)
	for i := start; i <= start+quantity; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func BenchmarkValidateSubsequence(b *testing.B) {
	numbers := generateNumbers(0, 12345)
	subSequence := []int{numbers[0], numbers[450], numbers[6000], numbers[9000], numbers[12000]}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidateSubsequence(numbers, subSequence)
	}
}
