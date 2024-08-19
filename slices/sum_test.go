package slices

import "testing"

func TestSum(t *testing.T) {

	type testCase[T int | int32 | int64 | float32 | float64 | string] struct {
		name     string
		values   []T
		expected T
	}

	testCasesForInt := []testCase[int]{
		{
			name:     "Collection of 5 values",
			values:   []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Collection with one value",
			values:   []int{1},
			expected: 1,
		},
		{
			name:     "Empty collection",
			values:   []int{},
			expected: 0,
		},
	}

	for _, tc := range testCasesForInt {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.values)
			if got != tc.expected {
				t.Errorf("got %d, want %d", got, tc.expected)
			}
		})
	}

	testCasesForFloat := []testCase[float64]{
		{
			name:     "Collection of 5 values",
			values:   []float64{1.2, 2.3, 3.4, 4.5, 5.6},
			expected: 17.0,
		},
		{
			name:     "Collection with one value",
			values:   []float64{1.2},
			expected: 1.2,
		},
		{
			name:     "Empty collection",
			values:   []float64{},
			expected: 0,
		},
	}

	for _, tc := range testCasesForFloat {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.values)
			if got != tc.expected {
				t.Errorf("got %f, want %f", got, tc.expected)
			}
		})
	}

	testCasesForString := []testCase[string]{
		{
			name:     "Collection of 5 strings",
			values:   []string{"the ", "quick ", "brown ", "fox ", "jumps"},
			expected: "the quick brown fox jumps",
		},
		{
			name:     "Collection with one string",
			values:   []string{"the"},
			expected: "the",
		},
		{
			name:     "Empty collection",
			values:   []string{},
			expected: "",
		},
	}

	for _, tc := range testCasesForString {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.values)
			if got != tc.expected {
				t.Errorf("got %s, want %s", got, tc.expected)
			}
		})
	}

}
