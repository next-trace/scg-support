# scg-support

A collection of utility packages for supply chain guard applications.

## Packages

See package documentation via:

- Local: ./scg docs (generates docs.txt and can serve godoc)
- Online: pkg.go.dev/github.com/next-trace/scg-support/util

### util

The `util` package provides a comprehensive collection of generic, type-safe utility functions for working with slices in Go. It complements Go's standard library `slices` package with additional functionality inspired by the ergonomics of Laravel's collection helpers.

All slice utility functions are accessible through a single import:

```go
import "github.com/next-trace/scg-support/util"
```

The package is organized into several logical groups of functions:

#### Core Functions
- **Map**: Transforms each element in a slice using a mapping function
- **Filter**: Creates a new slice with elements that pass a predicate function
- **Unique**: Removes duplicate values from a slice while preserving order
- **Pluck**: Extracts a specific property from a slice of structs
- **Chunk**: Splits a slice into chunks of a specified size
- **Flatten**: Transforms a slice of slices into a single flattened slice
- **GroupBy**: Groups slice elements by a key selector function
- **Reduce**: Reduces a slice to a single value using an accumulator
- **Intersect**: Returns elements common to all provided slices

#### Additional Functions
- **Contains**: Checks if a slice contains a specific element
- **IndexOf**: Returns the index of the first occurrence of an element
- **LastIndexOf**: Returns the index of the last occurrence of an element
- **Difference**: Returns elements in the first slice but not in other slices
- **Union**: Returns unique elements from all provided slices
- **ForEach**: Executes a function for each element in a slice
- **Reverse**: Returns a new slice with elements in reverse order
- **Take**: Returns the first n elements of a slice
- **Drop**: Returns a slice with the first n elements removed

#### Advanced Functions
- **MapReduce**: Combines Map and Reduce operations in a single pass
- **FindFirst**: Returns the first element that satisfies a predicate
- **FindLast**: Returns the last element that satisfies a predicate
- **Partition**: Divides a slice into two based on a predicate
- **Zip**: Combines elements from two slices into pairs
- **ZipWithIndex**: Pairs each element with its index
- **Shuffle**: Randomly reorders elements in a slice

## Development

### SCG Support Tool

The `scg` script is a bash utility tool that helps with various development tasks. It provides a convenient way to build, test, lint, and check security of the codebase.

To use the script, make sure it's executable:

```bash
chmod +x ./scg
```

#### Usage

```bash
./scg [command]
```

Available commands:

- `build` - Build the code using `go build -v ./...`
- `test` - Run tests with race detection and coverage using `go test -race -v -parallel 4 -coverprofile=coverage.txt -covermode=atomic ./...`
- `lint` - Run linter on the codebase using golangci-lint
- `lint-fix` - Run linter and fix issues automatically when possible
- `security` - Run security checks using:
  - `govulncheck` - Checks for known vulnerabilities in dependencies
  - `gosec` - Inspects source code for security problems
- `ci` - Run all CI checks (build, test, lint, security) and verifies Go version compatibility
- `install-tools` - Install required tools (golangci-lint, govulncheck, gosec)
- `help` - Show help message

Examples:

```bash
# Build the code
./scg build

# Run tests
./scg test

# Run linter and fix issues
./scg lint-fix

# Run all CI checks
./scg ci
```

The script will automatically install any missing tools required for the checks. When running the `ci` command, it will also check if you're using the recommended Go version (1.25) and warn you if you're using a different version.

### Quality Assurance

This project uses GitHub Actions for continuous integration and quality assurance:

- **Automated Testing**: All code changes are automatically tested
- **Linting**: Code quality is enforced using golangci-lint with a comprehensive set of linters
- **Security Scanning**: Dependencies are checked for vulnerabilities using govulncheck

The CI/CD pipeline runs on every push to the main branch and on pull requests. It ensures that all tests pass, the code meets quality standards, and there are no known security vulnerabilities in the dependencies.

#### Running CI Locally

You can run all CI checks locally using:

```bash
./scg ci
```

This command will:
1. Build the code
2. Run tests with race detection and coverage
3. Run the linter (golangci-lint)
4. Run security checks (govulncheck and gosec)

#### Usage Example

```go
package main

import (
    "fmt"
    "github.com/next-trace/scg-support/util"
)

func main() {
    // Map example
    numbers := []int{1, 2, 3, 4, 5}
    doubled := util.Map(numbers, func(n int, i int) int {
        return n * 2
    })
    fmt.Println(doubled) // [2 4 6 8 10]

    // Filter example
    evens := util.Filter(numbers, func(n int, i int) bool {
        return n%2 == 0
    })
    fmt.Println(evens) // [2 4]

    // Chunk example
    chunks := util.Chunk(numbers, 2)
    fmt.Println(chunks) // [[1 2] [3 4] [5]]

    // GroupBy example
    grouped := util.GroupBy(numbers, func(n int) string {
        if n%2 == 0 {
            return "even"
        }
        return "odd"
    })
    fmt.Println(grouped) // map[even:[2 4] odd:[1 3 5]]
}
```
