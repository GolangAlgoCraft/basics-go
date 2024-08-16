package slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("Should return the sum of all numbers in a slice", func(t *testing.T) {
		numbers := []int{4, 6, 32, 5, 3, 8, 10, 1, 2, 9}
		got := Sum(numbers)
		want := 80
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Should return 0 if an empty slice is provided", func(t *testing.T) {
		numbers := []int{}
		got := Sum(numbers)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
