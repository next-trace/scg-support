// Package util provides utility functions for working with slices.
// It complements Go's standard library slices package with additional functionality.
package util

import "slices"

// Contains checks if a slice contains a specific element.
// It returns true if the element is found, false otherwise.
//
// Note: For Go 1.21+, consider using slices.Contains from the standard library.
func Contains[S ~[]E, E comparable](collection S, element E) bool {
	// Use slices.Contains if available in your Go version
	return slices.Contains(collection, element)
}

// IndexOf returns the index of the first occurrence of an element in a slice.
// It returns -1 if the element is not found.
//
// Note: For Go 1.21+, consider using slices.Index from the standard library.
func IndexOf[S ~[]E, E comparable](collection S, element E) int {
	// Use slices.Index if available in your Go version
	return slices.Index(collection, element)
}

// LastIndexOf returns the index of the last occurrence of an element in a slice.
// It returns -1 if the element is not found.
func LastIndexOf[S ~[]E, E comparable](collection S, element E) int {
	for i := len(collection) - 1; i >= 0; i-- {
		if collection[i] == element {
			return i
		}
	}
	return -1
}

// Difference returns a new slice containing elements that are in the first slice
// but not in any of the other slices.
func Difference[S ~[]E, E comparable](first S, others ...S) S {
	if first == nil {
		return nil
	}

	// Create a map to track elements in other slices
	exclude := make(map[E]struct{})
	for _, other := range others {
		for _, item := range other {
			exclude[item] = struct{}{}
		}
	}

	var result S
	for _, item := range first {
		if _, found := exclude[item]; !found {
			result = append(result, item)
		}
	}
	return result
}

// Union returns a new slice containing unique elements from all provided slices.
// The order of elements is preserved based on their first occurrence across all slices.
func Union[S ~[]E, E comparable](slices ...S) S {
	if len(slices) == 0 {
		return nil
	}

	seen := make(map[E]struct{})
	var result S

	for _, slice := range slices {
		for _, item := range slice {
			if _, exists := seen[item]; !exists {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}

// ForEach executes a provided function once for each slice element.
func ForEach[S ~[]E, E any](collection S, action func(item E, index int)) {
	for i, item := range collection {
		action(item, i)
	}
}

// Reverse returns a new slice with the elements in reverse order.
//
// Note: For Go 1.21+, consider using slices.Clone and slices.Reverse from the standard library.
func Reverse[S ~[]E, E any](collection S) S {
	if collection == nil {
		return nil
	}

	length := len(collection)
	if length == 0 {
		return S{}
	}

	result := make(S, length)
	copy(result, collection)
	slices.Reverse(result)
	return result
}

// Take returns a new slice containing the first n elements of the original slice.
// If n is greater than the length of the slice, the entire slice is returned.
func Take[S ~[]E, E any](collection S, n int) S {
	if collection == nil {
		return nil
	}

	if n <= 0 {
		return S{}
	}

	length := len(collection)
	if n >= length {
		return slices.Clone(collection)
	}

	return slices.Clone(collection[:n])
}

// Drop returns a new slice with the first n elements removed.
// If n is greater than the length of the slice, an empty slice is returned.
func Drop[S ~[]E, E any](collection S, n int) S {
	if collection == nil {
		return nil
	}

	length := len(collection)
	if n <= 0 {
		return slices.Clone(collection)
	}

	if n >= length {
		return S{}
	}

	return slices.Clone(collection[n:])
}
