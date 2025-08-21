package chapter09_dynamic_programming

import (
	"sort"
	"time"
)

type event struct {
	name  string
	start time.Time
	end   time.Time
	value int
}

type taskNode struct {
	index              int
	event              event
	previousCompatible *taskNode
	maxValueUpToHere   int
}

type taskPlanner struct {
	nodes []taskNode
}

func planTasks(events []event) []string {
	if len(events) == 0 {
		return []string{}
	}

	planner := &taskPlanner{}
	planner.prepare(events)
	planner.solve()
	return planner.reconstructSolution()
}

func (tp *taskPlanner) prepare(events []event) {
	sort.Slice(events, func(i, j int) bool {
		return events[i].end.Before(events[j].end)
	})

	tp.nodes = make([]taskNode, len(events))
	for i, evt := range events {
		tp.nodes[i] = taskNode{
			event: evt,
			index: i,
		}
	}

	tp.findPreviousCompatible()
}

func (tp *taskPlanner) findPreviousCompatible() {
	for i := range tp.nodes {
		compatibleIndex := tp.binarySearchCompatible(i)
		if compatibleIndex != -1 {
			tp.nodes[i].previousCompatible = &tp.nodes[compatibleIndex]
		}
	}
}

func (tp *taskPlanner) binarySearchCompatible(eventIndex int) int {
	low, high := 0, eventIndex-1
	result := -1

	for low <= high {
		mid := (low + high) / 2
		if !tp.nodes[mid].event.end.After(tp.nodes[eventIndex].event.start) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

func (tp *taskPlanner) solve() {
	if len(tp.nodes) == 0 {
		return
	}

	tp.nodes[0].maxValueUpToHere = tp.nodes[0].event.value

	for i := 1; i < len(tp.nodes); i++ {
		includingCurrent := tp.nodes[i].event.value
		if tp.nodes[i].previousCompatible != nil {
			includingCurrent += tp.nodes[i].previousCompatible.maxValueUpToHere
		}

		excludingCurrent := tp.nodes[i-1].maxValueUpToHere
		if includingCurrent > excludingCurrent {
			tp.nodes[i].maxValueUpToHere = includingCurrent
		} else {
			tp.nodes[i].maxValueUpToHere = excludingCurrent
		}
	}
}

func (tp *taskPlanner) reconstructSolution() []string {
	if len(tp.nodes) == 0 {
		return []string{}
	}

	var result []string
	currentNode := &tp.nodes[len(tp.nodes)-1]

	for currentNode != nil {
		includingValue := currentNode.event.value
		if currentNode.previousCompatible != nil {
			includingValue += currentNode.previousCompatible.maxValueUpToHere
		}

		if includingValue == currentNode.maxValueUpToHere {
			result = append([]string{currentNode.event.name}, result...)
			currentNode = currentNode.previousCompatible
		} else {
			if currentNode.index > 0 {
				currentNode = &tp.nodes[currentNode.index-1]
			} else {
				break
			}
		}
	}

	return result
}
