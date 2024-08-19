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
		// int32 tests
		{name: "Sorted int32 slice", values: []int32{1, 2, 3, 4, 5}, expected: true},
		{name: "Unsorted int32 slice", values: []int32{1, 3, 2}, expected: false},
		{name: "Empty int32 slice", values: []int32{}, expected: true},
		{name: "Single element int32 slice", values: []int32{1}, expected: true},
		
		// int64 tests
		{name: "Sorted int64 slice", values: []int64{10, 20, 30, 40, 50}, expected: true},
		{name: "Unsorted int64 slice", values: []int64{10, 30, 20}, expected: false},
		{name: "Empty int64 slice", values: []int64{}, expected: true},
		{name: "Single element int64 slice", values: []int64{10}, expected: true},

		// float32 tests
		{name: "Sorted float32 slice", values: []float32{1.1, 2.2, 3.3}, expected: true},
		{name: "Unsorted float32 slice", values: []float32{1.1, 3.3, 2.2}, expected: false},
		{name: "Empty float32 slice", values: []float32{}, expected: true},
		{name: "Single element float32 slice", values: []float32{1.1}, expected: true},

		// float64 tests
		{name: "Sorted float64 slice", values: []float64{1.1, 2.2, 3.3}, expected: true},
		{name: "Unsorted float64 slice", values: []float64{1.1, 3.3, 2.2}, expected: false},
		{name: "Empty float64 slice", values: []float64{}, expected: true},
		{name: "Single element float64 slice", values: []float64{1.1}, expected: true},

		// string tests
		{name: "Sorted string slice", values: []string{"apple", "banana", "cherry"}, expected: true},
		{name: "Unsorted string slice", values: []string{"banana", "apple", "cherry"}, expected: false},
		{name: "Empty string slice", values: []string{}, expected: true},
		{name: "Single element string slice", values: []string{"apple"}, expected: true},
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
