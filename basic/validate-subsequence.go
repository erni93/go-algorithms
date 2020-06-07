package basic

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
