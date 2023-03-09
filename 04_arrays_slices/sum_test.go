package _4_arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(tb testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("non-empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 3, 6})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("slice with length 1", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{0, 3, 6})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

	t.Run("slice with length 0", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 3, 6})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
