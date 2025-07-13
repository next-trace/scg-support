// Package util provides generic, type-safe helpers that complement Go's standard
// library when working with slices. The goal of this package is to offer
// small, composable building blocks that are:
//
//   - Generic and domain-agnostic: utilities should be widely reusable across
//     services and not encode business logic.
//   - Predictable in semantics: nil-in, nil-out where appropriate; preserve
//     order unless otherwise stated; avoid mutating inputs.
//   - Ergonomic while remaining explicit: inspired by the convenience of
//     collection helpers from other ecosystems, but idiomatic Go.
//
// Design notes
//
//   - All functions are pure with respect to their inputs: any function that
//     returns a slice allocates a new slice (or returns nil/empty) and never
//     mutates its arguments.
//   - Functions follow consistent nil/empty semantics, documented on each API.
//     For example, Filter(nil, ...) returns nil, Chunk([]T{}, n) returns an
//     empty (non-nil) slice of chunks when n >= 1, and Shuffle returns a cloned
//     slice when length <= 1.
//   - Where possible, we leverage generics constraints to keep APIs type-safe
//     and avoid reflection.
//
// Quick start
//
//	import "github.com/next-trace/scg-support/util"
//
//	nums := []int{1, 2, 3, 4}
//	doubled := util.Map(nums, func(n int, _ int) int { return n * 2 })
//	evens := util.Filter(nums, func(n int, _ int) bool { return n%2 == 0 })
//
// # Testing and quality
//
// This repository is intentionally kept highly disciplined. Utilities must be
// generic and come with comprehensive table-driven tests. Use the provided
// ./scg script to run checks:
//
//	./scg ci
//
// That command builds, tests (with race and coverage), lints, and runs security
// checks. See README.md for more.
package util
