package _4_arrays_slices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
