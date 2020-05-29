package basic

// TwoNumbersSumFor returns two numbers of an array where its sum is equal to sum parameter, using double for
func TwoNumbersSumFor(numbers []int, sum int) []int {
	for i, num := range numbers {
		for _, num2 := range numbers[i:] {
			if num+num2 == sum {
				return []int{num, num2}
			}
		}
	}
	return []int{}
}

// TwoNumbersSumMap returns two numbers of an array where its sum is equal to sum parameter, using map
func TwoNumbersSumMap(numbers []int, sum int) []int {
	history := make(map[int]bool)
	for _, num := range numbers {
		if num <= sum {
			diff := sum - num
			if history[diff] {
				return []int{diff, num}
			}
			history[num] = true
		}
	}
	return []int{}
}

// TwoNumbersSumShrinking returns two numbers of an array where its sum is equal to sum parameter, shrinking from the corners.
// The numbers array must be shorted, low to high,
func TwoNumbersSumShrinking(numbers []int, sum int) []int {
	left := 0
	right := len(numbers) - 1
	for left != right {
		result := numbers[left] + numbers[right]
		if result == sum {
			return []int{numbers[left], numbers[right]}
		} else if result < sum {
			left++
		} else {
			right--
		}
	}
	return []int{}
}