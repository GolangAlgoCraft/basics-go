package array

import "errors"

type types interface {
	int | int32 | int64 | float32 | float64 | string
}

type Array[T types] struct {
	data   []T
	length int
}

// Creates a new array with the given values.
func New[T types](values ...T) *Array[T] {
	if len(values) > 0 {
		return &Array[T]{
			data:   values,
			length: len(values),
		}
	}

	return &Array[T]{
		data:   make([]T, 0),
		length: 0,
	}
}

// Adds the new value to the end of the array.
func (a *Array[T]) Push(value T) {
	a.data = append(a.data, value)
	a.length++
}

// Removes the last value from the array and returns it.
func (a *Array[T]) Pop() (T, error) {
	if a.length == 0 {
		var zero T
		return zero, errors.New("empty array")
	}

	value := a.data[a.length-1]
	a.data = a.data[:a.length-1]
	a.length--

	return value, nil
}

// Remove the value at a given index
func (a *Array[T]) DeleteAt(index int) error {
	if index < 0 || index >= a.length {
		return errors.New("index out of range")
	}
	a.data = append(a.data[:index], a.data[index+1:]...)
	a.length--
	return nil
}

// Insert a value at a given index
func (a *Array[T]) InsertAt(index int, value T) error {
	if index < 0 || index > a.length {
		return errors.New("index out of range")
	}
	a.data = append(a.data[:index], append([]T{value}, a.data[index:]...)...)
	a.length++
	return nil
}
