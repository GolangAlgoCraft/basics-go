package slices

import "testing"

func TestRemoveDuplicate(t *testing.T) {

	t.Run("Should remove duplicate values form a slice of strings", func(t *testing.T) {
		numbers := []string{"a", "b", "c", "a", "b", "d"}
		got := RemoveDuplicates(numbers)
		want := []string{"a", "b", "c", "d"}

		if len(got) != len(want) {
			t.Errorf("RemoveDuplicates() = %v, want %v", got, want)
		}

		for i, v := range got {
			if v != want[i] {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, want)
			}
		}
	})

	t.Run("Should remove duplicate values form a slice of integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 1, 2, 4}
		got := RemoveDuplicates(numbers)
		want := []int{1, 2, 3, 4}

		if len(got) != len(want) {
			t.Errorf("RemoveDuplicates() = %v, want %v", got, want)
		}

		for i, v := range got {
			if v != want[i] {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, want)
			}
		}
	})
}
