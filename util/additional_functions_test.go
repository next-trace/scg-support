package util

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("finds element in slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		if !Contains(input, 3) {
			t.Errorf("Contains() should have found 3 in %v", input)
		}
	})

	t.Run("returns false when element not in slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		if Contains(input, 6) {
			t.Errorf("Contains() should not have found 6 in %v", input)
		}
	})

	t.Run("works with strings", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		if !Contains(input, "banana") {
			t.Errorf("Contains() should have found 'banana' in %v", input)
		}
	})

	t.Run("returns false for empty slice", func(t *testing.T) {
		input := []int{}
		if Contains(input, 1) {
			t.Errorf("Contains() should return false for empty slice")
		}
	})

	t.Run("returns false for nil slice", func(t *testing.T) {
		var input []int
		if Contains(input, 1) {
			t.Errorf("Contains() should return false for nil slice")
		}
	})
}

func TestIndexOf(t *testing.T) {
	t.Run("finds index of element", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := 2
		result := IndexOf(input, 3)
		if result != expected {
			t.Errorf("IndexOf() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns -1 when element not found", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := -1
		result := IndexOf(input, 6)
		if result != expected {
			t.Errorf("IndexOf() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns first occurrence index", func(t *testing.T) {
		input := []int{1, 2, 3, 2, 1}
		expected := 1
		result := IndexOf(input, 2)
		if result != expected {
			t.Errorf("IndexOf() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns -1 for empty slice", func(t *testing.T) {
		input := []int{}
		expected := -1
		result := IndexOf(input, 1)
		if result != expected {
			t.Errorf("IndexOf() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns -1 for nil slice", func(t *testing.T) {
		var input []int
		expected := -1
		result := IndexOf(input, 1)
		if result != expected {
			t.Errorf("IndexOf() got = %v, want %v", result, expected)
		}
	})
}

func TestLastIndexOf(t *testing.T) {
	t.Run("finds last index of element", func(t *testing.T) {
		input := []int{1, 2, 3, 2, 1}
		expected := 3
		result := LastIndexOf(input, 2)
		if result != expected {
			t.Errorf("LastIndexOf() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns -1 when element not found", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := -1
		result := LastIndexOf(input, 6)
		if result != expected {
			t.Errorf("LastIndexOf() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns -1 for empty slice", func(t *testing.T) {
		input := []int{}
		expected := -1
		result := LastIndexOf(input, 1)
		if result != expected {
			t.Errorf("LastIndexOf() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns -1 for nil slice", func(t *testing.T) {
		var input []int
		expected := -1
		result := LastIndexOf(input, 1)
		if result != expected {
			t.Errorf("LastIndexOf() got = %v, want %v", result, expected)
		}
	})
}

func TestDifference(t *testing.T) {
	t.Run("returns elements in first slice but not in second", func(t *testing.T) {
		first := []int{1, 2, 3, 4, 5}
		second := []int{2, 4, 6}
		expected := []int{1, 3, 5}
		result := Difference(first, second)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Difference() got = %v, want %v", result, expected)
		}
	})

	t.Run("handles multiple other slices", func(t *testing.T) {
		first := []int{1, 2, 3, 4, 5}
		second := []int{2, 3}
		third := []int{4, 5}
		expected := []int{1}
		result := Difference(first, second, third)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Difference() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns all elements when no others provided", func(t *testing.T) {
		first := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		result := Difference(first)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Difference() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil for nil first slice", func(t *testing.T) {
		var first []int
		second := []int{1, 2, 3}
		result := Difference(first, second)
		if result != nil {
			t.Errorf("Difference() on nil first slice should return nil, but got %v", result)
		}
	})

	t.Run("handles nil other slices", func(t *testing.T) {
		first := []int{1, 2, 3}
		var second []int
		expected := []int{1, 2, 3}
		result := Difference(first, second)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Difference() got = %v, want %v", result, expected)
		}
	})
}

func TestUnion(t *testing.T) {
	t.Run("combines unique elements from all slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{2, 3, 4}
		slice3 := []int{3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		result := Union(slice1, slice2, slice3)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Union() got = %v, want %v", result, expected)
		}
	})

	t.Run("preserves order of first occurrence", func(t *testing.T) {
		slice1 := []int{3, 2, 1}
		slice2 := []int{4, 3, 2}
		expected := []int{3, 2, 1, 4}
		result := Union(slice1, slice2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Union() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns first slice when only one provided", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		result := Union(slice1)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Union() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil when no slices provided", func(t *testing.T) {
		result := Union[[]int]()
		if result != nil {
			t.Errorf("Union() with no slices should return nil, but got %v", result)
		}
	})

	t.Run("handles nil slices", func(t *testing.T) {
		var slice1 []int
		slice2 := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		result := Union(slice1, slice2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Union() got = %v, want %v", result, expected)
		}
	})
}

func TestForEach(t *testing.T) {
	t.Run("applies function to each element", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		sum := 0
		ForEach(input, func(item int, index int) {
			sum += item
		})
		expected := 15
		if sum != expected {
			t.Errorf("ForEach() sum got = %v, want %v", sum, expected)
		}
	})

	t.Run("provides correct indices", func(t *testing.T) {
		input := []int{10, 20, 30}
		indices := []int{}
		ForEach(input, func(_ int, index int) {
			indices = append(indices, index)
		})
		expected := []int{0, 1, 2}
		if !reflect.DeepEqual(indices, expected) {
			t.Errorf("ForEach() indices got = %v, want %v", indices, expected)
		}
	})

	t.Run("does nothing for empty slice", func(t *testing.T) {
		input := []int{}
		called := false
		ForEach(input, func(_ int, _ int) {
			called = true
		})
		if called {
			t.Errorf("ForEach() should not call function for empty slice")
		}
	})

	t.Run("does nothing for nil slice", func(t *testing.T) {
		var input []int
		called := false
		ForEach(input, func(_ int, _ int) {
			called = true
		})
		if called {
			t.Errorf("ForEach() should not call function for nil slice")
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("reverses elements in slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}
		result := Reverse(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Reverse() got = %v, want %v", result, expected)
		}
	})

	t.Run("works with strings", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		expected := []string{"c", "b", "a"}
		result := Reverse(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Reverse() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns empty slice for empty input", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := Reverse(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Reverse() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil for nil input", func(t *testing.T) {
		var input []int
		result := Reverse(input)
		if result != nil {
			t.Errorf("Reverse() on nil slice should return nil, but got %v", result)
		}
	})
}

func TestTake(t *testing.T) {
	t.Run("takes first n elements", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3}
		result := Take(input, 3)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Take() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns all elements when n >= length", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		result := Take(input, 5)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Take() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns empty slice when n <= 0", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []int{}
		result := Take(input, 0)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Take() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil for nil input", func(t *testing.T) {
		var input []int
		result := Take(input, 3)
		if result != nil {
			t.Errorf("Take() on nil slice should return nil, but got %v", result)
		}
	})
}

func TestDrop(t *testing.T) {
	t.Run("drops first n elements", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{4, 5}
		result := Drop(input, 3)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Drop() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns empty slice when n >= length", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []int{}
		result := Drop(input, 5)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Drop() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns all elements when n <= 0", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		result := Drop(input, 0)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Drop() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil for nil input", func(t *testing.T) {
		var input []int
		result := Drop(input, 3)
		if result != nil {
			t.Errorf("Drop() on nil slice should return nil, but got %v", result)
		}
	})
}
