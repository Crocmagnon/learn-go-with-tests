package _4_arrays_slices

import (
	"fmt"
	"reflect"
	"testing"
)

type sumCase struct {
	nums []int
	want int
}

func TestSum(t *testing.T) {
	cases := []sumCase{
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{2}, 2},
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3}, 6},
		{[]int{4, 6, 8}, 18},
		{[]int{4, 6, 8, 12}, 30},
	}

	for _, c := range cases {
		t.Run(testName(c), func(t *testing.T) {
			got := Sum(c.nums)

			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}

		})
	}
}

func testName(c sumCase) string {
	return fmt.Sprintf("%v -> %d", c.nums, c.want)
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
