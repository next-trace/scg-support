package util

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("converts slice of ints to slice of strings", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []string{"1", "2", "3"}
		result := Map(input, func(item int, _ int) string {
			return strconv.Itoa(item)
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Map() got = %v, want %v", result, expected)
		}
	})

	t.Run("handles nil slice", func(t *testing.T) {
		var input []int
		result := Map(input, func(item int, _ int) int { return item })
		if result != nil {
			t.Errorf("Map() on a nil slice should return nil, but got %v", result)
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("filters for even numbers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		expected := []int{2, 4, 6}
		result := Filter(input, func(item int, _ int) bool {
			return item%2 == 0
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Filter() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns empty slice when no items match", func(t *testing.T) {
		input := []int{1, 3, 5}
		result := Filter(input, func(item int, index int) bool { return item%2 == 0 })
		if len(result) != 0 {
			t.Errorf("Filter() should have returned an empty slice, but got %v", result)
		}
	})

	t.Run("returns nil for nil slice", func(t *testing.T) {
		var input []int
		result := Filter(input, func(item int, _ int) bool { return true })
		if result != nil {
			t.Errorf("Filter() on nil slice should return nil, got %v", result)
		}
	})
}

func TestUnique(t *testing.T) {
	t.Run("removes duplicates and preserves order", func(t *testing.T) {
		input := []string{"a", "b", "a", "c", "b", "d", "a"}
		expected := []string{"a", "b", "c", "d"}
		result := Unique(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique() got = %v, want %v", result, expected)
		}
	})

	t.Run("works on integers", func(t *testing.T) {
		input := []int{1, 1, 2, 5, 2, 3, 1, 5}
		expected := []int{1, 2, 5, 3}
		result := Unique(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Unique() with ints got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil for nil slice", func(t *testing.T) {
		var input []int
		result := Unique(input)
		if result != nil {
			t.Errorf("Unique() on nil slice should return nil, got %v", result)
		}
	})

	t.Run("returns nil for empty slice", func(t *testing.T) {
		input := []int{}
		result := Unique(input)
		if result != nil {
			t.Errorf("Unique() on empty slice should return nil, got %v", result)
		}
	})
}

func TestPluck(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}
	t.Run("plucks a string property from a struct slice", func(t *testing.T) {
		input := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		}
		expected := []string{"Alice", "Bob", "Charlie"}
		result := Pluck(input, func(item User) string {
			return item.Name
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Pluck() for Name got = %v, want %v", result, expected)
		}
	})

	t.Run("plucks an int property from a struct slice", func(t *testing.T) {
		input := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		}
		expected := []int{1, 2, 3}
		result := Pluck(input, func(item User) int {
			return item.ID
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Pluck() for ID got = %v, want %v", result, expected)
		}
	})

	t.Run("returns empty slice for empty input", func(t *testing.T) {
		var input []User // actually nil leads to nil return; test both cases
		resultNil := Pluck(input, func(item User) int { return item.ID })
		if resultNil != nil {
			t.Errorf("Pluck() on nil slice should return nil, got %v", resultNil)
		}
		empty := []User{}
		resultEmpty := Pluck(empty, func(item User) int { return item.ID })
		if resultEmpty == nil || len(resultEmpty) != 0 {
			t.Errorf("Pluck() on empty slice should return empty (len 0) non-nil slice, got %v", resultEmpty)
		}
	})
}

func TestChunk(t *testing.T) {
	t.Run("chunks a slice into equal parts", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
		result := Chunk(input, 2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Chunk() got = %v, want %v", result, expected)
		}
	})

	t.Run("handles uneven chunks", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := [][]int{{1, 2}, {3, 4}, {5}}
		result := Chunk(input, 2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Chunk() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns empty slice for empty input", func(t *testing.T) {
		input := []int{}
		expected := [][]int{}
		result := Chunk(input, 2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Chunk() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil for nil input", func(t *testing.T) {
		var input []int
		result := Chunk(input, 2)
		if result != nil {
			t.Errorf("Chunk() on nil slice should return nil, but got %v", result)
		}
	})

	t.Run("returns nil for invalid size", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Chunk(input, 0)
		if result != nil {
			t.Errorf("Chunk() with size 0 should return nil, but got %v", result)
		}
	})
}

func TestFlatten(t *testing.T) {
	t.Run("flattens a slice of slices", func(t *testing.T) {
		input := [][]int{{1, 2}, {3, 4}, {5, 6}}
		expected := []int{1, 2, 3, 4, 5, 6}
		result := Flatten(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Flatten() got = %v, want %v", result, expected)
		}
	})

	t.Run("handles empty inner slices", func(t *testing.T) {
		input := [][]int{{1, 2}, {}, {5, 6}}
		expected := []int{1, 2, 5, 6}
		result := Flatten(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Flatten() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns empty slice for empty input", func(t *testing.T) {
		input := [][]int{}
		expected := []int{}
		result := Flatten(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Flatten() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil for nil input", func(t *testing.T) {
		var input [][]int
		result := Flatten(input)
		if result != nil {
			t.Errorf("Flatten() on nil slice should return nil, but got %v", result)
		}
	})
}

func TestGroupBy(t *testing.T) {
	t.Run("groups integers by even/odd", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		expected := map[string][]int{
			"even": {2, 4, 6},
			"odd":  {1, 3, 5},
		}
		result := GroupBy(input, func(item int) string {
			if item%2 == 0 {
				return "even"
			}
			return "odd"
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() got = %v, want %v", result, expected)
		}
	})

	t.Run("groups strings by length", func(t *testing.T) {
		input := []string{"a", "bb", "ccc", "d", "ee"}
		expected := map[int][]string{
			1: {"a", "d"},
			2: {"bb", "ee"},
			3: {"ccc"},
		}
		result := GroupBy(input, func(item string) int {
			return len(item)
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns empty map for empty input", func(t *testing.T) {
		input := []int{}
		expected := map[bool][]int{}
		result := GroupBy(input, func(item int) bool {
			return item%2 == 0
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil for nil input", func(t *testing.T) {
		var input []int
		result := GroupBy(input, func(item int) bool {
			return item%2 == 0
		})
		if result != nil {
			t.Errorf("GroupBy() on nil slice should return nil, but got %v", result)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("sums integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := 15
		result := Reduce(input, 0, func(acc int, item int, index int) int {
			return acc + item
		})
		if result != expected {
			t.Errorf("Reduce() got = %v, want %v", result, expected)
		}
	})

	t.Run("concatenates strings", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		expected := "abc"
		result := Reduce(input, "", func(acc string, item string, index int) string {
			return acc + item
		})
		if result != expected {
			t.Errorf("Reduce() got = %v, want %v", result, expected)
		}
	})

	t.Run("uses index in reducer", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := 10 // Sum of indices: 0+1+2+3+4 = 10
		result := Reduce(input, 0, func(acc int, item int, index int) int {
			return acc + index
		})
		if result != expected {
			t.Errorf("Reduce() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns initial value for empty input", func(t *testing.T) {
		input := []int{}
		expected := 10
		result := Reduce(input, 10, func(acc int, item int, index int) int {
			return acc + item
		})
		if result != expected {
			t.Errorf("Reduce() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns initial value for nil input", func(t *testing.T) {
		var input []int
		expected := 10
		result := Reduce(input, 10, func(acc int, item int, index int) int {
			return acc + item
		})
		if result != expected {
			t.Errorf("Reduce() got = %v, want %v", result, expected)
		}
	})
}

func TestIntersect(t *testing.T) {
	t.Run("finds common elements", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4}
		slice2 := []int{3, 4, 5, 6}
		slice3 := []int{3, 4, 7, 8}
		expected := []int{3, 4}
		result := Intersect(slice1, slice2, slice3)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Intersect() got = %v, want %v", result, expected)
		}
	})

	t.Run("preserves order from first slice", func(t *testing.T) {
		slice1 := []int{4, 3, 2, 1}
		slice2 := []int{1, 2, 3, 4}
		expected := []int{4, 3, 2, 1}
		result := Intersect(slice1, slice2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Intersect() got = %v, want %v", result, expected)
		}
	})

	t.Run("handles duplicates in first slice", func(t *testing.T) {
		slice1 := []int{1, 2, 2, 3, 3, 4}
		slice2 := []int{2, 3, 4}
		expected := []int{2, 3, 4}
		result := Intersect(slice1, slice2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Intersect() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns empty slice when no common elements", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{3, 4}
		expected := []int{}
		result := Intersect(slice1, slice2)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Intersect() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns first slice when only one slice provided", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		result := Intersect(slice1)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Intersect() got = %v, want %v", result, expected)
		}
	})

	t.Run("returns nil when no slices provided", func(t *testing.T) {
		result := Intersect[[]int]()
		if result != nil {
			t.Errorf("Intersect() with no slices should return nil, but got %v", result)
		}
	})
}
