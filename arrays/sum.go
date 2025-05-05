package array

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAll(numbers ...[]int) []int {
	sumAll := func(acc, x []int) []int {
		return append(acc, Sum(x))
	}

	return Reduce(numbers, sumAll, []int{})
}

// The tail of a collection is all items in the collection except the first one (the "head").
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbers, sumTail, []int{})
}

func Reduce[T any](collection []T, f func(T, T) T, initialValue T) T {
	var result = initialValue
	for _, el := range collection {
		result = f(result, el)
	}
	return result
}
