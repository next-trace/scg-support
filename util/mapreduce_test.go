package util

import (
	"reflect"
	"strings"
	"testing"
)

type assertErr struct{}

func (assertErr) Error() string { return "forced error" }

func TestMapReduce(t *testing.T) {
	t.Run("sums squares of integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := 55 // 1² + 2² + 3² + 4² + 5² = 55
		result := MapReduce(
			input,
			func(item int, _ int) int { return item * item },
			0,
			func(acc int, mapped int, _ int) int { return acc + mapped },
		)
		if result != expected {
			t.Errorf("MapReduce() got = %v, want %v", result, expected)
		}
	})

	t.Run("builds a comma-separated string", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		expected := "APPLE, BANANA, CHERRY"
		result := MapReduce(
			input,
			func(item string, _ int) string { return strings.ToUpper(item) },
			"",
			func(acc string, mapped string, index int) string {
				if index == 0 {
					return mapped
				}
				return acc + ", " + mapped
			},
		)
		if result != expected {
			t.Errorf("MapReduce() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns initial value for empty input", func(t *testing.T) {
		input := []int{}
		expected := 10
		result := MapReduce(
			input,
			func(item int, _ int) int { return item },
			10,
			func(acc int, mapped int, _ int) int { return acc + mapped },
		)
		if result != expected {
			t.Errorf("MapReduce() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns initial value for nil input", func(t *testing.T) {
		var input []int
		expected := 10
		result := MapReduce(
			input,
			func(item int, _ int) int { return item },
			10,
			func(acc int, mapped int, _ int) int { return acc + mapped },
		)
		if result != expected {
			t.Errorf("MapReduce() got = %v, want %v", result, expected)
		}
	})
}

func TestFindFirst(t *testing.T) {
	t.Run("finds first even number", func(t *testing.T) {
		input := []int{1, 3, 4, 6, 7, 8}
		expected := 4
		result, found := FindFirst(input, func(item int, _ int) bool {
			return item%2 == 0
		})
		if !found || result != expected {
			t.Errorf("FindFirst() got = (%v, %v), want = (%v, true)", result, found, expected)
		}
	})

	t.Run("returns false when no match", func(t *testing.T) {
		input := []int{1, 3, 5, 7, 9}
		_, found := FindFirst(input, func(item int, _ int) bool {
			return item%2 == 0
		})
		if found {
			t.Errorf("FindFirst() should return found=false when no match")
		}
	})

	t.Run("returns false for empty slice", func(t *testing.T) {
		input := []int{}
		_, found := FindFirst(input, func(item int, _ int) bool {
			return true
		})
		if found {
			t.Errorf("FindFirst() should return found=false for empty slice")
		}
	})

	t.Run("returns false for nil slice", func(t *testing.T) {
		var input []int
		_, found := FindFirst(input, func(item int, _ int) bool {
			return true
		})
		if found {
			t.Errorf("FindFirst() should return found=false for nil slice")
		}
	})
}

func TestFindLast(t *testing.T) {
	t.Run("finds last even number", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7}
		expected := 6
		result, found := FindLast(input, func(item int, _ int) bool {
			return item%2 == 0
		})
		if !found || result != expected {
			t.Errorf("FindLast() got = (%v, %v), want = (%v, true)", result, found, expected)
		}
	})

	t.Run("returns false when no match", func(t *testing.T) {
		input := []int{1, 3, 5, 7, 9}
		_, found := FindLast(input, func(item int, _ int) bool {
			return item%2 == 0
		})
		if found {
			t.Errorf("FindLast() should return found=false when no match")
		}
	})

	t.Run("returns false for empty slice", func(t *testing.T) {
		input := []int{}
		_, found := FindLast(input, func(item int, _ int) bool {
			return true
		})
		if found {
			t.Errorf("FindLast() should return found=false for empty slice")
		}
	})

	t.Run("returns false for nil slice", func(t *testing.T) {
		var input []int
		_, found := FindLast(input, func(item int, _ int) bool {
			return true
		})
		if found {
			t.Errorf("FindLast() should return found=false for nil slice")
		}
	})
}

func TestPartition(t *testing.T) {
	t.Run("partitions even and odd numbers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		expectedEven := []int{2, 4, 6}
		expectedOdd := []int{1, 3, 5}
		even, odd := Partition(input, func(item int, _ int) bool {
			return item%2 == 0
		})
		if !reflect.DeepEqual(even, expectedEven) || !reflect.DeepEqual(odd, expectedOdd) {
			t.Errorf("Partition() got = (%v, %v), want = (%v, %v)", even, odd, expectedEven, expectedOdd)
		}
	})

	t.Run("handles all items matching", func(t *testing.T) {
		input := []int{2, 4, 6, 8}
		expectedMatched := []int{2, 4, 6, 8}
		expectedUnmatched := []int{}
		matched, unmatched := Partition(input, func(item int, index int) bool {
			return item%2 == 0
		})
		if !reflect.DeepEqual(matched, expectedMatched) || !reflect.DeepEqual(unmatched, expectedUnmatched) {
			t.Errorf("Partition() got = (%v, %v), want = (%v, %v)", matched, unmatched, expectedMatched, expectedUnmatched)
		}
	})

	t.Run("handles no items matching", func(t *testing.T) {
		input := []int{1, 3, 5, 7}
		expectedMatched := []int{}
		expectedUnmatched := []int{1, 3, 5, 7}
		matched, unmatched := Partition(input, func(item int, index int) bool {
			return item%2 == 0
		})
		if !reflect.DeepEqual(matched, expectedMatched) || !reflect.DeepEqual(unmatched, expectedUnmatched) {
			t.Errorf("Partition() got = (%v, %v), want = (%v, %v)", matched, unmatched, expectedMatched, expectedUnmatched)
		}
	})

	t.Run("returns nil, nil for nil input", func(t *testing.T) {
		var input []int
		matched, unmatched := Partition(input, func(item int, index int) bool {
			return true
		})
		if matched != nil || unmatched != nil {
			t.Errorf("Partition() on nil slice should return (nil, nil), but got (%v, %v)", matched, unmatched)
		}
	})
}

func TestZip(t *testing.T) {
	t.Run("zips two slices of same length", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []string{"a", "b", "c"}
		expected := [][2]any{
			{1, "a"},
			{2, "b"},
			{3, "c"},
		}
		result := Zip(slice1, slice2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Zip() got = %v, want %v", result, expected)
		}
	})

	t.Run("zips to the length of the shorter slice", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := []string{"a", "b", "c"}
		expected := [][2]any{
			{1, "a"},
			{2, "b"},
			{3, "c"},
		}
		result := Zip(slice1, slice2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Zip() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil when first slice is nil", func(t *testing.T) {
		var slice1 []int
		slice2 := []string{"a", "b", "c"}
		result := Zip(slice1, slice2)
		if result != nil {
			t.Errorf("Zip() with nil first slice should return nil, but got %v", result)
		}
	})

	t.Run("returns nil when second slice is nil", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		var slice2 []string
		result := Zip(slice1, slice2)
		if result != nil {
			t.Errorf("Zip() with nil second slice should return nil, but got %v", result)
		}
	})

	t.Run("returns empty slice when one input is empty", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []string{"a", "b"}
		result := Zip(slice1, slice2)
		if result == nil || len(result) != 0 {
			t.Errorf("Zip() with empty first slice should return empty (len 0) non-nil slice, got %v", result)
		}
	})

	t.Run("returns empty slice when both inputs are empty", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []string{}
		result := Zip(slice1, slice2)
		if result == nil || len(result) != 0 {
			t.Errorf("Zip() with both empty slices should return empty (len 0) non-nil slice, got %v", result)
		}
	})
}

func TestZipWithIndex(t *testing.T) {
	t.Run("pairs elements with their indices", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		expected := [][2]any{
			{"a", 0},
			{"b", 1},
			{"c", 2},
		}
		result := ZipWithIndex(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ZipWithIndex() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil for nil input", func(t *testing.T) {
		var input []string
		result := ZipWithIndex(input)
		if result != nil {
			t.Errorf("ZipWithIndex() on nil slice should return nil, but got %v", result)
		}
	})

	t.Run("returns empty slice for empty input", func(t *testing.T) {
		input := []int{}
		expected := [][2]any{}
		result := ZipWithIndex(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ZipWithIndex() got = %v, want %v", result, expected)
		}
	})
}

func TestShuffle(t *testing.T) {
	// Save and restore readRandom for test isolation
	origReadRandom := readRandom
	t.Cleanup(func() { readRandom = origReadRandom })
	t.Run("returns a shuffled copy", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := Shuffle(input)

		// Check that result has the same length
		if len(result) != len(input) {
			t.Errorf("Shuffle() returned slice of length %v, want %v", len(result), len(input))
		}

		// Check that result contains all the same elements
		inputMap := make(map[int]int)
		resultMap := make(map[int]int)

		for _, v := range input {
			inputMap[v]++
		}

		for _, v := range result {
			resultMap[v]++
		}

		if !reflect.DeepEqual(inputMap, resultMap) {
			t.Errorf("Shuffle() result contains different elements than input")
		}

		// Check that the order is different (this could occasionally fail by random chance)
		// We'll consider it shuffled if at least one element is in a different position
		different := false
		for i, v := range input {
			if i < len(result) && v != result[i] {
				different = true
				break
			}
		}

		// Only fail if we have enough elements and they're all in the same order
		if len(input) > 3 && !different {
			t.Errorf("Shuffle() did not change the order of elements")
		}
	})

	t.Run("returns nil for nil input", func(t *testing.T) {
		var input []int
		result := Shuffle(input)
		if result != nil {
			t.Errorf("Shuffle() on nil slice should return nil, but got %v", result)
		}
	})

	t.Run("handles single element slice", func(t *testing.T) {
		input := []int{1}
		expected := []int{1}
		result := Shuffle(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Shuffle() got = %v, want %v", result, expected)
		}
	})

	t.Run("handles empty slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := Shuffle(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Shuffle() got = %v, want %v", result, expected)
		}
	})

	t.Run("covers two-byte random path (len > 256)", func(t *testing.T) {
		// length 300 ensures i will exceed 255 in the loop
		input := make([]int, 300)
		for i := range input {
			input[i] = i
		}
		result := Shuffle(input)
		if len(result) != len(input) {
			t.Errorf("Shuffle() two-byte path length mismatch: got %d want %d", len(result), len(input))
		}
	})

	t.Run("covers four-byte random path (len > 65536)", func(t *testing.T) {
		// length 70000 ensures i will exceed 65535 in the loop
		input := make([]int, 70000)
		for i := range input {
			input[i] = i
		}
		result := Shuffle(input)
		if len(result) != len(input) {
			t.Errorf("Shuffle() four-byte path length mismatch: got %d want %d", len(result), len(input))
		}
	})

	t.Run("returns unshuffled clone on random error", func(t *testing.T) {
		// Force readRandom to return error
		readRandom = func(b []byte) (int, error) { return 0, assertErr{} }
		input := []int{1, 2, 3, 4, 5}
		result := Shuffle(input)
		// Order should be unchanged
		if !reflect.DeepEqual(result, input) {
			t.Errorf("Shuffle() on error should return unshuffled clone; got %v want %v", result, input)
		}
		// And result should be a different underlying array than input
		if len(result) > 0 {
			result[0] = 999
			if input[0] == 999 {
				t.Errorf("Shuffle() on error should return a clone, not alias the input slice")
			}
		}
	})
}
