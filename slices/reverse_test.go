package slices

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		values   interface{}
		expected interface{}
	}{
		// test case for int
		{name: "slice of 5 int", values: []int{1, 2, 3, 4, 5}, expected: []int{5, 4, 3, 2, 1}},
		{name: "slice of one int", values: []int{1}, expected: []int{1}},
		{name: "Empty int slice", values: []int{}, expected: []int{}},
		// test case for float64
		{name: "slice of 5 float64", values: []float64{1.2, 2.3, 3.4, 4.5, 5.6}, expected: []float64{5.6, 4.5, 3.4, 2.3, 1.2}},
		{name: "slice of one float64", values: []float64{1.2}, expected: []float64{1.2}},
		{name: "Empty float64 slice", values: []float64{}, expected: []float64{}},
		// test case for string
		{name: "slice of 5 string", values: []string{"the ", "quick ", "brown ", "fox ", "jumps"}, expected: []string{"jumps", "fox ", "brown ", "quick ", "the "}},
		{name: "slice of one string", values: []string{"the"}, expected: []string{"the"}},
		{name: "Empty string slice", values: []string{}, expected: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got interface{}
			var err error
			switch v := tt.values.(type) {
			case []int:
				got = Reverse(v)
				err = compare[int](got.([]int), tt.expected.([]int))
			case []float64:
				got = Reverse(v)
				err = compare[float64](got.([]float64), tt.expected.([]float64))
			case []string:
				got = Reverse(v)
				err = compare[string](got.([]string), tt.expected.([]string))
			default:
				t.Fatalf("unsupported type %T", v)
			}

			if err != nil {
				t.Error(err.Error())
			}

		})
	}

}

func compare[T int | string | float64](got, expected []T) error {
	if len(got) != len(expected) {
		return fmt.Errorf("Reverse() = %v, want %v", got, expected)
	}

	for i, v := range got {
		if v != expected[i] {
			return fmt.Errorf("Reverse() = %v, want %v", got, expected)
		}
	}
	return nil
}
