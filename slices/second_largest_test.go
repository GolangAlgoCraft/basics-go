package slices

import (
	"testing"
)

func TestSecondLargest(t *testing.T) {
	t.Run("Should return the second largest from a given slice", func(t *testing.T) {
		initialSlice := []int{1,9,2,3,4,5}
		got := SecondLargest(initialSlice)
		want := 5

		// compare the 2 values
	
			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
	})
}
