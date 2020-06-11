package basic

import "fmt"

func ValidateSubsequence(numbers []int, subsequence []int) bool {
	if len(numbers) < len(subsequence) {
		return false
	}
	seqIndex := 0
	for _, number := range numbers {
		if number == subsequence[seqIndex] {
			if seqIndex == len(subsequence)-1 {
				return true
			} else {
				seqIndex++
			}
		}
	}
	return false
}

func DemoValidateSubsequence() {
	fmt.Printf("[ValidateSubsequence] Numbers: %v Subsequence: %v -> %v\n", []int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}, ValidateSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}))
}
