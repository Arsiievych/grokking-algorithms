package chapter06_breadth_first_search

import (
	"reflect"
	"testing"
)

func TestFindShorterWay(t *testing.T) {
	type testCase struct {
		name    string
		from    string
		to      string
		want    []string
		wantAny [][]string
	}

	cases := []testCase{
		{
			name: "Returns_same_point_when_from_equals_to",
			from: "Херсон",
			to:   "Херсон",
			want: []string{"Херсон"},
		},
		{
			name: "Returns_direct_route_when_direct_connection_exists",
			from: "Луцьк",
			to:   "Рівне",
			want: []string{"Луцьк", "Рівне"},
		},
		{
			name: "Returns_empty_path_if_destination_missing_in_graph",
			from: "Київ",
			to:   "New York",
			want: []string{},
		},
		{
			name: "Returns_empty_path_if_source_missing_in_graph",
			from: "Berlin",
			to:   "Київ",
			want: []string{},
		},
		{
			name: "Finds_shortest_route_through_multiple_cities",
			from: "Кропивницький",
			to:   "Львів",
			want: []string{"Кропивницький", "Вінниця", "Житомир", "Рівне", "Львів"},
		},
		{
			name: "Returns_any_of_multiple_shortest_routes",
			from: "Луцьк",
			to:   "Тернопіль",
			wantAny: [][]string{
				{"Луцьк", "Рівне", "Тернопіль"},
				{"Луцьк", "Львів", "Тернопіль"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindShorterWay(tc.from, tc.to)

			if tc.wantAny != nil {
				matched := false
				for _, option := range tc.wantAny {
					if reflect.DeepEqual(got, option) {
						matched = true
						break
					}
				}
				if !matched {
					t.Errorf("got %v, expected one of %v", got, tc.wantAny)
				}
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
