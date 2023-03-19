package _4_arrays_slices

func Sum(numbers []int) int {
	return Reduce[int](numbers, 0, func(a, b int) int { return a + b })
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	sumTail := func(acc, item []int) []int {
		if len(item) == 0 {
			return append(acc, 0)
		} else {
			tail := item[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbersToSum, []int{}, sumTail)
}

type accumulator[T any] func(T, T) T

func Reduce[T any](collection []T, initialValue T, accumulate accumulator[T]) T {
	result := initialValue
	for _, item := range collection {
		result = accumulate(result, item)
	}
	return result
}
