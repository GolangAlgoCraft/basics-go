package array

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("create empty array", func(t *testing.T) {
		array := New[int]()
		if array == nil {
			t.Fatal("Expected array to be created, got nil")
		}
		if array.length != 0 {
			t.Errorf("Expected length 0, got %d", array.length)
		}
	})

	t.Run("creates a new array with initial values", func(t *testing.T) {
		// create a new array with some initial values
		array := New(1, 2, 3)

		// check if the array is created
		if array == nil {
			t.Fatal("Expected array to be created, got nil")
		}

		// check the length
		if array.length != 3 {
			t.Errorf("Expected length 3, got %d", array.length)
		}
	})

}

func TestPush(t *testing.T) {
	t.Run("pushing on empty array", func(t *testing.T) {
		array := New[int]()
		array.Push(1)
		if array.length != 1 {
			t.Errorf("Expected length 1, got %d", array.length)
		}
		if array.data[0] != 1 {
			t.Errorf("Expected value 1, got %d", array.data[0])
		}

	})

	t.Run("pushing to array with initialized values", func(t *testing.T) {
		array := New[int](3, 6, 2, 7)
		array.Push(1)
		if array.length != 5 {
			t.Errorf("Expected length 1, got %d", array.length)
		}
		if array.data[4] != 1 {
			t.Errorf("Expected value 1, got %d", array.data[0])
		}

	})
}

func TestPop(t *testing.T) {
	t.Run("popping from empty array", func(t *testing.T) {
		array := New[int]()
		_, err := array.Pop()
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("popping from array with initialized values", func(t *testing.T) {
		array := New[int](3, 6, 2, 7)
		value, err := array.Pop()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if value != 7 {
			t.Errorf("Expected value 7, got %d", value)
		}
		if array.length != 3 {
			t.Errorf("Expected length 3, got %d", array.length)
		}
	})
}

func TestDeleteAt(t *testing.T) {
	t.Run("deleting from empty array", func(t *testing.T) {
		array := New[int]()
		err := array.DeleteAt(0)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("deleting from array with initialized values", func(t *testing.T) {
		array := New[int](3, 6, 2, 7)
		err := array.DeleteAt(1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if array.length != 3 {
			t.Errorf("Expected length 3, got %d", array.length)
		}
		if array.data[1] != 2 {
			t.Errorf("Expected value 2, got %d", array.data[1])
		}
	})
}

func TestInsertAt(t *testing.T) {
	t.Run("inserting into empty array", func(t *testing.T) {
		array := New[int]()
		err := array.InsertAt(0, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if array.length != 1 {
			t.Errorf("Expected length 1, got %d", array.length)
		}
		if array.data[0] != 1 {
			t.Errorf("Expected value 1, got %d", array.data[0])
		}
	})

	t.Run("using out of range index", func(t *testing.T) {
		array := New[int](3, 6, 2, 7)
		err := array.InsertAt(5, 1)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run(("inserting into array with initialized values"), func(t *testing.T) {
		array := New[int](3, 6, 2, 7)
		err := array.InsertAt(1, 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if array.length != 5 {
			t.Errorf("Expected length 5, got %d", array.length)
		}
		if array.data[1] != 1 {
			t.Errorf("Expected value 1, got %d", array.data[1])
		}
	})
}
