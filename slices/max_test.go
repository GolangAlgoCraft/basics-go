package slices

import "testing"

func TestMax(t *testing.T) {
	t.Run("Should return the maximum value of a slice", func(t *testing.T) {
		numbers := []int{4, 6, 32, 5, 3, 8, 10, 1, 2, 9}
		got := Max(numbers)
		want := 32
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Should return 0 if empty slice is provided", func(t *testing.T) {
		numbers := []int{}
		got := Max(numbers)
		want := 0
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
