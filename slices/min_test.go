package slices

import "testing"

func TestMin(t *testing.T) {
	t.Run("Should return the minimum value of a slice", func(t *testing.T) {
		numbers := []int{9, 6, 2, 7, 4, 1, 5, 8, 3}
		got := Min(numbers)
		want := 1
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Should return 0 if empty slice is provided", func(t *testing.T) {
		numbers := []int{}
		got := Min(numbers)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
