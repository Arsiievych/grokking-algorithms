# Grokking Algorithms

A collection of algorithmic solutions and concepts implemented in Go, based on "Grokking Algorithms" by Aditya Bhargava.

## Structure

The project is organized by chapters, with each chapter focusing on a specific algorithm or concept:

- **Chapter 1**: Binary Search
- **Chapter 2**: Selection Sort
- **Chapter 3**: Recursion vs Loop

Each chapter directory contains implementation files (`.go`) and corresponding test files (`_test.go`).

## Running Tests

This project uses Go's standard testing package for both basic tests and benchmarks.

### Running Basic Tests

To run all tests in the project:

```bash
go test ./...
```

To run tests for a specific chapter:

```bash
go test ./chapter01-binary_search
go test ./chapter02-selection_sort
go test ./chapter03-recursion
```

To run a specific test:

```bash
go test ./chapter01-binary_search -run TestBinarySearch
```

### Running Benchmark Tests

To run all benchmarks in the project:

```bash
go test -bench=. ./...
```

To run benchmarks for a specific chapter:

```bash
go test -bench=. ./chapter01-binary_search
go test -bench=. ./chapter02-selection_sort
go test -bench=. ./chapter03-recursion
```

To run a specific benchmark:

```bash
go test -bench=BenchmarkBinarySearch ./chapter01-binary_search
```

### Benchmark Output Example

When running benchmarks, you'll see output similar to this:

```
goos: darwin
goarch: amd64
pkg: github.com/Arsiievych/grokking-algorithms/chapter03-recursion
BenchmarkSumLoop-8             1000000              1234 ns/op
BenchmarkSumRecursion-8         500000              2345 ns/op
BenchmarkGaussFormula-8       10000000               123 ns/op
PASS
ok      github.com/Arsiievych/grokking-algorithms/chapter03-recursion   3.456s
```

This output shows:
- The operating system and architecture
- The package being tested
- For each benchmark:
  - The name of the benchmark and number of CPU cores used
  - The number of iterations performed
  - The average time per operation in nanoseconds
- The total time taken to run all benchmarks

### Additional Test Options

- `-v`: Verbose output, shows all tests being run
- `-count=N`: Run each test N times
- `-timeout=T`: Set a timeout for the test, e.g., `-timeout=30s`
- `-benchtime=T`: Run each benchmark for duration T (e.g., `-benchtime=5s`)
- `-benchmem`: Include memory allocation statistics in benchmark results
