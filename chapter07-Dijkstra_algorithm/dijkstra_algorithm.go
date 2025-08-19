package chapter07_Dijkstra_algorithm

import (
	"container/heap"
	"math"
)

func Dijkstra(graph map[string]map[string]int, start, target string) []string {
	if start == target {
		return []string{start}
	}

	if !isValidPath(graph, start, target) {
		return nil
	}

	distances, parents := initializeDijkstra(graph, start)
	visited := make(map[string]bool)

	pq := createInitialPriorityQueue(start)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item).id
		if visited[current] {
			continue
		}
		visited[current] = true

		if current == target {
			return buildPath(parents, start, target)
		}

		relaxNeighbors(graph, current, distances, parents, pq, visited)
	}

	return nil
}

func isValidPath(graph map[string]map[string]int, start, target string) bool {
	_, startExists := graph[start]
	_, targetExists := graph[target]
	return startExists && targetExists
}

func initializeDijkstra(graph map[string]map[string]int, start string) (map[string]int, map[string]string) {
	distances := make(map[string]int)
	parents := make(map[string]string)

	for node := range graph {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	return distances, parents
}

func createInitialPriorityQueue(start string) *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{
		id:       start,
		priority: 0,
	})
	return pq
}

func relaxNeighbors(graph map[string]map[string]int, current string,
	distances map[string]int, parents map[string]string,
	pq *PriorityQueue, visited map[string]bool) {

	for neighbor, weight := range graph[current] {
		if visited[neighbor] {
			continue
		}

		newDistance := distances[current] + weight
		if newDistance < distances[neighbor] {
			distances[neighbor] = newDistance
			parents[neighbor] = current

			heap.Push(pq, &Item{
				id:       neighbor,
				priority: newDistance,
			})
		}
	}
}

func buildPath(parents map[string]string, start, target string) []string {
	path := []string{}
	current := target

	for current != "" {
		path = append([]string{current}, path...)
		if current == start {
			break
		}
		current = parents[current]
	}

	return path
}
