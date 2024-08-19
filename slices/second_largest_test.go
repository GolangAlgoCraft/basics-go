package slices

import "testing"

func TestSecondLargest(t *testing.T) {
	tests := []struct {
		name     string
		values   interface{}
		expected interface{}
	}{
		// int32 test cases
		{name: "Second largest in int32 slice", values: []int32{1, 3, 2}, expected: int32(2)},
		{name: "All elements are the same in int32 slice", values: []int32{2, 2, 2}, expected: int32(0)},
		{name: "Single element in int32 slice", values: []int32{5}, expected: int32(0)},
		{name: "Two elements in int32 slice", values: []int32{5, 10}, expected: int32(5)},
		
		// int64 test cases
		{name: "Second largest in int64 slice", values: []int64{10, 20, 30, 40, 50}, expected: int64(40)},
		{name: "All elements are the same in int64 slice", values: []int64{100, 100, 100}, expected: int64(0)},
		{name: "Single element in int64 slice", values: []int64{15}, expected: int64(0)},
		{name: "Two elements in int64 slice", values: []int64{15, 30}, expected: int64(15)},
		
		// float32 test cases
		{name: "Second largest in float32 slice", values: []float32{1.1, 3.3, 2.2}, expected: float32(2.2)},
		{name: "All elements are the same in float32 slice", values: []float32{1.1, 1.1, 1.1}, expected: float32(0)},
		{name: "Single element in float32 slice", values: []float32{5.5}, expected: float32(0)},
		{name: "Two elements in float32 slice", values: []float32{5.5, 10.5}, expected: float32(5.5)},
		
		// float64 test cases
		{name: "Second largest in float64 slice", values: []float64{5.5, 10.5, 7.7}, expected: float64(7.7)},
		{name: "All elements are the same in float64 slice", values: []float64{10.5, 10.5, 10.5}, expected: float64(0)},
		{name: "Single element in float64 slice", values: []float64{12.3}, expected: float64(0)},
		{name: "Two elements in float64 slice", values: []float64{7.7, 10.5}, expected: float64(7.7)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got interface{}
			switch v := tt.values.(type) {
			case []int32:
				got = SecondLargest(v)
			case []int64:
				got = SecondLargest(v)
			case []float32:
				got = SecondLargest(v)
			case []float64:
				got = SecondLargest(v)
			default:
				t.Fatalf("unsupported type %T", v)
			}

			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
