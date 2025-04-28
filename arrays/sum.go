package array

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return result
}
