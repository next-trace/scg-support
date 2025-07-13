// Package util provides generic, type-safe utility functions for working with slices,
// inspired by the ergonomics of Laravel's helpers. It complements Go's standard library
// `slices` package with additional functionality.
package util

// Map applies a function to each element of a slice, returning a new slice
// containing the results. It is a type-safe Go equivalent of Laravel's `Arr::map`.
func Map[S ~[]E, E, R any](collection S, iteratee func(item E, index int) R) []R {
	if collection == nil {
		return nil
	}

	result := make([]R, len(collection))
	for index, item := range collection {
		result[index] = iteratee(item, index)
	}
	return result
}

// Filter iterates over elements of a slice, returning a new slice containing all elements
// for which the predicate function returns true. This is the Go equivalent of `Arr::where`.
func Filter[S ~[]E, E any](collection S, predicate func(item E, index int) bool) S {
	if collection == nil {
		return nil
	}

	// The zero value for a slice is nil, which is often desired.
	var result S
	for index, item := range collection {
		if predicate(item, index) {
			result = append(result, item)
		}
	}
	return result
}

// Unique returns a new slice with duplicate values removed.
// The order of elements is preserved from the first time they appear in the collection.
// It requires the element type to be comparable.
func Unique[S ~[]E, E comparable](collection S) S {
	if collection == nil {
		return nil
	}

	seen := make(map[E]struct{}, len(collection))

	var result S
	for _, item := range collection {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Pluck creates a slice of a single property from a slice of structs or maps.
// It is a type-safe Go equivalent of Laravel's `Arr::pluck`.
func Pluck[S ~[]E, E, R any](collection S, propertyGetter func(item E) R) []R {
	if collection == nil {
		return nil
	}
	result := make([]R, len(collection))
	for index, item := range collection {
		result[index] = propertyGetter(item)
	}
	return result
}

// Chunk splits a slice into chunks of the specified size.
// The last chunk may contain fewer elements if the slice length is not divisible by size.
// If size is less than 1, it returns nil.
func Chunk[S ~[]E, E any](collection S, size int) []S {
	if collection == nil || size < 1 {
		return nil
	}

	length := len(collection)
	if length == 0 {
		return []S{}
	}

	chunksCount := (length + size - 1) / size // Ceiling division
	chunks := make([]S, 0, chunksCount)

	for i := 0; i < length; i += size {
		end := i + size
		if end > length {
			end = length
		}
		chunks = append(chunks, collection[i:end])
	}

	return chunks
}

// Flatten transforms a slice of slices into a single flattened slice.
func Flatten[E any](collections [][]E) []E {
	if collections == nil {
		return nil
	}

	// Calculate total length to avoid reallocations
	totalLen := 0
	for _, collection := range collections {
		totalLen += len(collection)
	}

	result := make([]E, 0, totalLen)
	for _, collection := range collections {
		result = append(result, collection...)
	}
	return result
}

// GroupBy groups the elements of a slice by the result of the keySelector function.
// It returns a map where each key is the result of the keySelector function and
// the value is a slice of all elements that produced that key.
func GroupBy[S ~[]E, E any, K comparable](collection S, keySelector func(item E) K) map[K]S {
	if collection == nil {
		return nil
	}

	result := make(map[K]S)
	for _, item := range collection {
		key := keySelector(item)
		result[key] = append(result[key], item)
	}
	return result
}

// Reduce applies a function against an accumulator and each element in the slice
// to reduce it to a single value.
func Reduce[S ~[]E, E, R any](collection S, initialValue R, reducer func(acc R, item E, index int) R) R {
	if len(collection) == 0 {
		return initialValue
	}

	result := initialValue
	for i, item := range collection {
		result = reducer(result, item, i)
	}
	return result
}

// Intersect returns a slice containing all elements that are present in all given slices.
// The order of elements is preserved from the first slice.
func Intersect[S ~[]E, E comparable](collections ...S) S {
	if len(collections) == 0 {
		return nil
	}

	if len(collections) == 1 {
		return collections[0]
	}

	// Create a map to count occurrences of each element
	elementCounts := make(map[E]int)

	// Count occurrences in all slices except the first one
	for _, collection := range collections[1:] {
		// Use a set to avoid counting duplicates within the same slice
		seen := make(map[E]bool)
		for _, item := range collection {
			if !seen[item] {
				elementCounts[item]++
				seen[item] = true
			}
		}
	}

	// Filter the first slice to include only elements that appear in all other slices
	var result S
	seen := make(map[E]bool) // To handle duplicates in the first slice

	for _, item := range collections[0] {
		if !seen[item] && elementCounts[item] == len(collections)-1 {
			result = append(result, item)
			seen[item] = true
		}
	}

	// Return an empty slice (not nil) if no common elements were found
	if len(result) == 0 && len(collections[0]) > 0 {
		return S{}
	}

	return result
}
