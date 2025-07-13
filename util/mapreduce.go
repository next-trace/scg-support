// Package util provides utility functions for working with slices.
package util

import (
	"crypto/rand"
	"encoding/binary"
	"slices"
)

// readRandom is an indirection for crypto/rand.Read to enable testing error paths.
var readRandom = rand.Read

// MapReduce combines Map and Reduce operations in a single pass.
// It applies a mapping function to each element of a slice and then reduces the results
// to a single value using a reducer function.
func MapReduce[S ~[]E, E, M, R any](
	collection S,
	mapper func(item E, index int) M,
	initialValue R,
	reducer func(acc R, mapped M, index int) R,
) R {
	if len(collection) == 0 {
		return initialValue
	}

	result := initialValue
	for i, item := range collection {
		mapped := mapper(item, i)
		result = reducer(result, mapped, i)
	}
	return result
}

// FindFirst returns the first element in a slice that satisfies a predicate function.
// It returns the found element and a boolean indicating whether an element was found.
func FindFirst[S ~[]E, E any](collection S, predicate func(item E, index int) bool) (E, bool) {
	var zero E
	if len(collection) == 0 {
		return zero, false
	}

	for i, item := range collection {
		if predicate(item, i) {
			return item, true
		}
	}
	return zero, false
}

// FindLast returns the last element in a slice that satisfies a predicate function.
// It returns the found element and a boolean indicating whether an element was found.
func FindLast[S ~[]E, E any](collection S, predicate func(item E, index int) bool) (E, bool) {
	var zero E
	if len(collection) == 0 {
		return zero, false
	}

	for i := len(collection) - 1; i >= 0; i-- {
		item := collection[i]
		if predicate(item, i) {
			return item, true
		}
	}
	return zero, false
}

// Partition divides a slice into two slices based on a predicate function.
// The first returned slice contains all elements that satisfy the predicate,
// and the second contains all elements that don't.
func Partition[S ~[]E, E any](collection S, predicate func(item E, index int) bool) (S, S) {
	if collection == nil {
		return nil, nil
	}

	var matched, unmatched S
	for i, item := range collection {
		if predicate(item, i) {
			matched = append(matched, item)
		} else {
			unmatched = append(unmatched, item)
		}
	}

	// Ensure we return empty slices (not nil) when no items match or all items match
	if len(matched) == 0 && len(collection) > 0 {
		matched = S{}
	}
	if len(unmatched) == 0 && len(collection) > 0 {
		unmatched = S{}
	}

	return matched, unmatched
}

// Zip combines elements from two slices into a slice of pairs.
// The length of the result is the minimum of the lengths of the two input slices.
// Each pair is represented as a [2]any array where the first element is from the first slice
// and the second element is from the second slice.
func Zip[S1 ~[]E1, E1 any, S2 ~[]E2, E2 any](slice1 S1, slice2 S2) [][2]any {
	if slice1 == nil || slice2 == nil {
		return nil
	}

	len1 := len(slice1)
	len2 := len(slice2)
	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	if minLen == 0 {
		return [][2]any{}
	}

	result := make([][2]any, minLen)
	for i := range result {
		result[i] = [2]any{slice1[i], slice2[i]}
	}
	return result
}

// ZipWithIndex pairs each element in a slice with its index.
// Each pair is represented as a [2]any array where the first element is the original element
// and the second element is its index.
func ZipWithIndex[S ~[]E, E any](collection S) [][2]any {
	if collection == nil {
		return nil
	}

	length := len(collection)
	if length == 0 {
		return [][2]any{}
	}

	result := make([][2]any, length)
	for i, item := range collection {
		result[i] = [2]any{item, i}
	}
	return result
}

// Shuffle returns a new slice with the elements randomly reordered.
// It uses crypto/rand for secure random number generation.
//
// This function uses a cryptographically secure random number generator
// and is suitable for both general-purpose and security-sensitive operations.
func Shuffle[S ~[]E, E any](collection S) S {
	if collection == nil {
		return nil
	}

	length := len(collection)
	if length <= 1 {
		return slices.Clone(collection)
	}

	// Create a copy to avoid modifying the original
	result := slices.Clone(collection)

	// Fisher-Yates shuffle algorithm with crypto/rand
	for i := length - 1; i > 0; i-- {
		// Generate a random number in the range [0, i]
		// We only need enough random bytes to cover the range [0, i]
		maxBytes := 1
		if i > 255 {
			maxBytes = 2 // 2 bytes for i > 255
		}
		if i > 65535 {
			maxBytes = 4 // 4 bytes for i > 65535
		}

		randomBytes := make([]byte, maxBytes)
		_, err := readRandom(randomBytes)
		if err != nil {
			// In case of error, return the unshuffled clone
			return result
		}

		// Convert bytes to an integer and reduce to the range [0, i]
		var randomInt int
		switch maxBytes {
		case 1:
			randomInt = int(randomBytes[0]) % (i + 1)
		case 2:
			randomInt = int(binary.BigEndian.Uint16(randomBytes)) % (i + 1)
		case 4:
			// This is safe because we're only using 4 bytes (uint32) which fits in int on all platforms
			randomInt = int(binary.BigEndian.Uint32(randomBytes)) % (i + 1)
		}

		j := randomInt

		// Swap elements
		result[i], result[j] = result[j], result[i]
	}

	return result
}
