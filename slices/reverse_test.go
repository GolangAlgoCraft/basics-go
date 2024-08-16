package slices

import (
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("Should reverse a given slice", func(t *testing.T) {
		initialSlice := []int{1, 2, 3, 4, 5}
		reversedSlice := Reverse(initialSlice)
		want := []int{5, 4, 3, 2, 1}

		// compare the 2 slices
		for i, v := range reversedSlice {
			if v != want[i] {
				t.Errorf("got %d want %d", reversedSlice, want)
			}
		}
	})
}
