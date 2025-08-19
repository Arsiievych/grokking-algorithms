package chapter07_Dijkstra_algorithm

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	type testCase = struct {
		name   string
		graph  map[string]map[string]int
		start  string
		target string
		want   []string
	}

	cases := []testCase{
		{
			name:   "start and target are the same",
			start:  "A",
			target: "A",
			want:   []string{"A"},
		},
		{
			name: "start missing in graph",
			graph: map[string]map[string]int{
				"B": {"A": 1},
			},
			start:  "A",
			target: "B",
			want:   nil,
		},
		{
			name: "target missing in graph",
			graph: map[string]map[string]int{
				"A": {"B": 1},
			},
			start:  "A",
			target: "B",
			want:   nil,
		},
		{
			name: "multiple paths - shortest path selected\n",
			graph: map[string]map[string]int{
				"Start":  {"A": 5, "B": 2},
				"A":      {"D": 4, "C": 2},
				"B":      {"A": 8, "C": 7},
				"C":      {"Finish": 1},
				"D":      {"C": 6, "Finish": 3},
				"Finish": {},
			},
			start:  "Start",
			target: "Finish",
			want:   []string{"Start", "A", "C", "Finish"},
		},
		{
			name: "negative cycle avoided",
			graph: map[string]map[string]int{
				"Start":  {"A": 10},
				"A":      {"C": 20},
				"B":      {"A": 1},
				"C":      {"Finish": 30, "B": 1},
				"Finish": {},
			},
			start:  "Start",
			target: "Finish",
			want:   []string{"Start", "A", "C", "Finish"},
		},
		{
			name: "disconnected graph (no path)",
			graph: map[string]map[string]int{
				"A": {"C": 1},
				"B": {},
				"C": {},
			},
			start:  "A",
			target: "B",
			want:   nil,
		},
		{
			name: "single edge graph",
			graph: map[string]map[string]int{
				"A": {"B": 3},
				"B": {},
			},
			start:  "A",
			target: "B",
			want:   []string{"A", "B"},
		},
		{
			name: "zero-weight edge preferred",
			graph: map[string]map[string]int{
				"S": {"A": 0, "B": 5},
				"A": {"B": 1},
				"B": {},
			},
			start:  "S",
			target: "B",
			want:   []string{"S", "A", "B"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Dijkstra(tc.graph, tc.start, tc.target)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}

}
