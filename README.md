# Grokking Algorithms

A collection of algorithmic solutions and concepts implemented in Go, based on "Grokking Algorithms" by Aditya Bhargava.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Implemented Tasks](#implemented-tasks)
  - [Chapter 1: Binary Search](#chapter-1-binary-search)
  - [Chapter 2: Selection Sort](#chapter-2-selection-sort)
  - [Chapter 3: Recursion vs Loop](#chapter-3-recursion-vs-loop)
  - [Chapter 4: Quick Sort](#chapter-4-quick-sort)
  - [Chapter 6: Breadth-First Search (BFS)](#chapter-6-breadth-first-search-bfs)
  - [Chapter 7: Dijkstra's Algorithm](#chapter-7-dijkstras-algorithm)
  - [Chapter 8: Greedy Algorithm — Coin Change](#chapter-8-greedy-algorithm--coin-change)
  - [Chapter 9: Dynamic Programming — Task Planner](#chapter-9-dynamic-programming--task-planner)
- [Running Tests](#running-tests)
- [Additional Notes](#additional-notes)

## Prerequisites

- Go 1.24+ (see go.mod)

## Project Structure

Each chapter directory contains implementation files (`.go`) and corresponding test files (`_test.go`). The chapters follow the book’s progression and keep APIs minimal and idiomatic for Go.

## Implemented Tasks

### Chapter 1: Binary Search
Finds the index of a target value in a sorted slice in O(log n) time using a classic binary search. Returns -1 when the value is not present.

### Chapter 2: Selection Sort
Sorts a slice in ascending order by repeatedly selecting the minimum element. Simple, easy to follow, O(n²) time.

### Chapter 3: Recursion vs Loop
Compares iterative and recursive approaches (e.g., summation) and includes a constant-time variant (Gauss formula) to highlight trade-offs and performance differences.

### Chapter 4: Quick Sort
Implements a divide-and-conquer quicksort. Demonstrates partitioning and recursion for average O(n log n) performance.

### Chapter 6: Breadth-First Search (BFS)
Traverses a graph level by level to find the shortest path in number of edges or to locate a node that satisfies a predicate.

### Chapter 7: Dijkstra's Algorithm
Computes single-source shortest paths on directed weighted graphs using a simple priority queue.

### Chapter 8: Greedy Algorithm — Coin Change
Builds a representation of a given amount using available denominations (UAH/kopeck set used in tests). Shows how greedy strategies perform with canonical coin systems.

### Chapter 9: Dynamic Programming — Task Planner
Weighted Interval Scheduling: selects a set of non-overlapping events with maximum total value.
- Compatibility rule: an earlier event is compatible if it ends at or before the next starts (end <= start), allowing back-to-back scheduling.
- Approach: sort by end time, binary-search the previous compatible event, and use DP to compute the optimal total.
- Deterministic reconstruction: when include/exclude totals tie, the implementation prefers inclusion.
- Zero-duration events (start == end) are supported; negative-value events are naturally excluded if harmful.

## Running Tests

Run all tests in the repository:

```bash
go test ./...
```

Run tests for a specific chapter:

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

Run a specific test by name (example):

```bash
go test ./chapter01-binary_search -run TestBinarySearch
```

Or run the dynamic programming tests for the task planner only:

```bash
go test ./chapter09-dynamic_programming -run TestPlanTasks
```

## Additional Notes

- Chapter 8 (Greedy Coin Change) uses Ukrainian hryvnia/kopeck denominations in tests and outputs, for example: `5грн.`, `1грн.`, `50к.`, `25к.`, `10к.`, `5к.`, `2к.`, `1к.`
