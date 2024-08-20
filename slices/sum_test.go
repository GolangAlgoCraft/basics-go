package slices

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		values   interface{}
		expected interface{}
	}{
		// test case int
		{name: "Collection of 5 values", values: []int{1, 2, 3, 4, 5}, expected: 15},
		{name: "Collection with one value", values: []int{1}, expected: 1},
		{name: "Empty collection", values: []int{}, expected: 0},
		// test case float64
		{name: "Collection of 5 values", values: []float64{1.2, 2.3, 3.4, 4.5, 5.6}, expected: 17.0},
		{name: "Collection with one value", values: []float64{1.2}, expected: 1.2},
		{name: "Empty collection", values: []float64{}, expected: 0.0},
		// test case for string
		{name: "Collection of 5 values", values: []string{"the ", "quick ", "brown ", "fox ", "jumps"}, expected: "the quick brown fox jumps"},
		{name: "Collection with one value", values: []string{"the"}, expected: "the"},
		{name: "Empty collection", values: []string{}, expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got interface{}
			switch v := tt.values.(type) {
			case []int:
				got = Sum(v)
			case []float64:
				got = Sum(v)
			case []string:
				got = Sum(v)
			default:
				t.Fatalf("unsupported type %T", v)
			}

			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}

		})
	}

}
