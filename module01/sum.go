package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	sum := 0

	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}
