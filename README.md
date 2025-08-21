# Grokking Algorithms

A collection of algorithmic solutions and concepts implemented in Go, based on "Grokking Algorithms" by Aditya Bhargava.

## Prerequisites

- Go 1.24+ (see go.mod)

## Structure

The project is organized by chapters, with each chapter focusing on a specific algorithm or concept:

- **Chapter 1**: Binary Search
- **Chapter 2**: Selection Sort
- **Chapter 3**: Recursion vs Loop
- **Chapter 4**: Quick Sort
- **Chapter 6**: Breadth-First Search (BFS)
- **Chapter 7**: Dijkstra's Algorithm (with a simple priority queue)
- **Chapter 8**: Greedy Algorithm — Coin Change (UAH denominations)
- **Chapter 9**: Dynamic Programming — Task Planner (Weighted Interval Scheduling)

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
go test ./chapter04-quick_sort
go test ./chapter06-breadth_first_search
go test ./chapter07-Dijkstra_algorithm
go test ./chapter08-greedy_algorithm
go test ./chapter09-dynamic_programming
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
go test -bench=. ./chapter04-quick_sort
go test -bench=. ./chapter06-breadth_first_search
go test -bench=. ./chapter07-Dijkstra_algorithm
go test -bench=. ./chapter08-greedy_algorithm
go test -bench=. ./chapter09-dynamic_programming
```

To run a specific benchmark:

```bash
go test -bench=BenchmarkBinarySearch ./chapter01-binary_search
go test -bench=BenchmarkQuickSort ./chapter04-quick_sort
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

## Chapter 9: Dynamic Programming — Task Planner (Weighted Interval Scheduling)

This chapter implements the classic Weighted Interval Scheduling problem. The goal is to choose a set of non-overlapping events with maximum total value.

Key points:
- Events are considered compatible if an earlier event ends at or before the next one starts (end <= start). This allows back-to-back scheduling.
- The algorithm sorts by end time, finds the previous compatible event via binary search, and uses dynamic programming to compute the optimal value in O(n log n) time and O(n) space.
- Reconstruction prefers including the current event when the "include" and "exclude" totals are equal, making the choice deterministic.
- Zero-duration events (start == end) are supported and can fit between other events when compatible.
- Negative-value events are naturally excluded by DP if they would reduce the total value.

Run tests for this chapter:

```bash
go test ./chapter09-dynamic_programming -run TestPlanTasks
```

## Notes

- Chapter 8 (Greedy Coin Change) uses Ukrainian hryvnia/kopeck denominations in tests and outputs, for example: `5грн.`, `1грн.`, `50к.`, `25к.`, `10к.`, `5к.`, `2к.`, `1к.`
