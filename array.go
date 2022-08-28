package array

// Map creates a new slice with the results of calling a provided
// function on every element in given slice.
func Map[T1, T2 any](slice []T1, f func(T1) T2) []T2 {
	newSlice := make([]T2, 0, len(slice))
	for _, item := range slice {
		newSlice = append(newSlice, f(item))
	}
	return newSlice
}

// Filter creates a new slice with only elements that passes the
// condition inside the provided function.
func Filter[T any](slice []T, f func(T) bool) []T {
	newSlice := []T{}
	for _, item := range slice {
		if f(item) {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

// FilterErrors creates a new slice with elements returned by
// provided callback function. If the function returns non-nil error,
// FilterError skips the current element.
func FilterErrors[T1, T2 any](slice []T1, f func(T1) (T2, error)) []T2 {
	newSlice := []T2{}

	for _, element := range slice {
		if newElement, err := f(element); err == nil {
			newSlice = append(newSlice, newElement)
		}
	}

	return newSlice
}

// ForEach helps to loop over slice by executing a provided callback
// function for each element in a slice.
func ForEach[T any](slice []T, f func(T)) {
	for _, item := range slice {
		f(item)
	}
}

// Every checks every element in the slice that passes the condition,
// returning true or false as appropriate.
func Every[T any](slice []T, f func(T) bool) bool {
	if slice == nil {
		return false
	}

	for _, item := range slice {
		if !f(item) {
			return false
		}
	}

	return true
}

// Some checks if at least one element in the slice that passes the
// condition, returning true or false as appropriate.
func Some[T any](slice []T, f func(T) bool) bool {
	for _, item := range slice {
		if f(item) {
			return true
		}
	}
	return false
}

// Index returns the index of the first occurrence of the specified
// element in the slice, or -1 if it is not found.
func Index[T comparable](slice []T, item T) int {
	for index, sliceItem := range slice {
		if sliceItem == item {
			return index
		}
	}
	return -1
}

// IndexFunc returns the index of the first element in a slice that
// pass the test in a testing function.
func IndexFunc[T any](slice []T, f func(T) bool) int {
	for index, sliceItem := range slice {
		if f(sliceItem) {
			return index
		}
	}
	return -1
}

// Contains checks if a slice includes the element that passes the
// condition, returning true or false as appropriate.
func Contains[T comparable](slice []T, item T) bool {
	return Index(slice, item) >= 0
}

// Reduce applies a function against an accumulator and each element
// in the slice to reduce it to a single value.
func Reduce[T1, T2 any](slice []T1, initial T2, f func(accumulator T2, current T1) T2) T2 {
	acc := initial
	for _, curr := range slice {
		acc = f(acc, curr)
	}
	return acc
}

// Find returns the value of the first element in a slice that pass
// the test in a testing function; it also returns a bool saying
// whether the target is really found in the slice.
func Find[T any](slice []T, f func(T) bool) (T, bool) {
	for _, sliceItem := range slice {
		if f(sliceItem) {
			return sliceItem, true
		}
	}
	var zero T
	return zero, false
}

// Fill fills the elements in a slice with a static value and returns
// the modified slice. Fill returns nil if given slice == nil.
func Fill[T any](slice []T, value T, start, end int) []T {
	for index := range slice {
		if index >= start && index <= end {
			slice[index] = value
		}
	}
	return slice
}

// Reverse reverses a slice in place. Element at last index will be
// first and element at 0 index will be last. Reverse returns nil if
// given slice == nil.
func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// MapKeys returns slice of keys of given map.
func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
	newSlice := make([]Key, 0, len(m))
	for key := range m {
		newSlice = append(newSlice, key)
	}
	return newSlice
}

// MapKeys returns slice of values of given map.
func MapValues[Key comparable, Val any](m map[Key]Val) []Val {
	newSlice := make([]Val, 0, len(m))
	for _, val := range m {
		newSlice = append(newSlice, val)
	}
	return newSlice
}

// Range returns a slice of T values in range [start : start + n].
// if start <= 0, Range returns a slice of T values in range [:n]
// If start >= len(list), Range returns an empty slice.
// If n <= 0, Range returns a slice of T values in range [start:]
// If n > 0, Range returns at most n T elements.
func Range[T any](slice []T, start int, n int) []T {
	length := len(slice)

	if start <= 0 {
		if n > 0 && n < length {
			return slice[:n]
		}

		return slice
	}

	if start > 0 && start < length {
		if n > 0 && start+n < length {
			return slice[start : start+n]
		}

		if n > 0 && start+n >= length {
			return slice[start:]
		}

		if n <= 0 {
			return slice[start:]
		}
	}

	return []T{}
}
