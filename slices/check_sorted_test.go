package slices

import "testing"

type Test struct {
	name     string
	values   interface{}
	expected bool
}
// for more information about generic params using interface{} see: https://go.dev/doc/tutorial/generics
func TestCheckSorted(t *testing.T) {
	tests := []Test{
		{name: "Sorted int32 slice", values: []int32{1, 2, 3, 4, 5}, expected: true},
		{name: "Unsorted int32 slice", values: []int32{1, 3, 2}, expected: false},
		{name: "Sorted float64 slice", values: []float64{1.1, 2.2, 3.3}, expected: true},
		{name: "Sorted string slice", values: []string{"apple", "banana", "cherry"}, expected: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// dynamic type assertion
			var got bool
			switch v := tc.values.(type) {
			case []int32:
				got = CheckSorted(v)
			case []float64:
				got = CheckSorted(v)
			case []string:
				got = CheckSorted(v)
			default:
				t.Fatalf("unsupported type %T", v)
			}

			if got != tc.expected {
				t.Errorf("got %t, want %t", got, tc.expected)
			}
		})
	}
}
