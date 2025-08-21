package chapter09_dynamic_programming

import (
	"reflect"
	"testing"
	"time"
)

func TestPlanTasks(t *testing.T) {
	type testCase struct {
		name   string
		events []event
		want   []string
	}

	testCases := []testCase{
		{
			name:   "empty case",
			events: []event{},
			want:   []string{},
		},
		{
			name: "single event",
			events: []event{
				{name: "Meeting", start: parseHM("10:00"), end: parseHM("11:00"), value: 100},
			},
			want: []string{"Meeting"},
		},
		{
			name: "overlapping events with higher value",
			events: []event{
				{name: "Study", start: parseHM("09:00"), end: parseHM("12:00"), value: 200},
				{name: "Date", start: parseHM("10:30"), end: parseHM("11:15"), value: 400},
				{name: "Exercise", start: parseHM("11:00"), end: parseHM("13:00"), value: 300},
				{name: "Lunch", start: parseHM("13:00"), end: parseHM("14:00"), value: 100},
			},
			want: []string{"Date", "Lunch"},
		},
		{
			name: "full day schedule",
			events: []event{
				{name: "Sleep", start: parseHM("00:00"), end: parseHM("07:00"), value: 800},
				{name: "Breakfast", start: parseHM("07:30"), end: parseHM("08:30"), value: 100},
				{name: "Work", start: parseHM("09:00"), end: parseHM("17:00"), value: 500},
				{name: "Dinner", start: parseHM("18:00"), end: parseHM("19:00"), value: 150},
				{name: "Reading", start: parseHM("20:00"), end: parseHM("22:00"), value: 200},
				{name: "Sleep", start: parseHM("22:30"), end: parseHM("23:59"), value: 400},
			},
			want: []string{"Sleep", "Breakfast", "Work", "Dinner", "Reading", "Sleep"},
		},
		{
			name: "touching intervals are compatible",
			events: []event{
				{name: "B", start: parseHM("10:00"), end: parseHM("11:01"), value: 1},
				{name: "A", start: parseHM("09:00"), end: parseHM("10:00"), value: 1},
			},
			want: []string{"A", "B"},
		},
		{
			name: "all overlap choose max",
			events: []event{
				{name: "X", start: parseHM("09:00"), end: parseHM("12:00"), value: 100},
				{name: "Y", start: parseHM("10:00"), end: parseHM("13:00"), value: 200},
				{name: "Z", start: parseHM("11:00"), end: parseHM("14:00"), value: 150},
			},
			want: []string{"Y"},
		},
		{
			name: "many small beat big",
			events: []event{
				{name: "Big", start: parseHM("09:00"), end: parseHM("17:00"), value: 300},
				{name: "S1", start: parseHM("09:00"), end: parseHM("11:00"), value: 150},
				{name: "S2", start: parseHM("11:00"), end: parseHM("13:00"), value: 150},
				{name: "S3", start: parseHM("13:00"), end: parseHM("17:00"), value: 150},
			},
			want: []string{"S1", "S2", "S3"},
		},
		{
			name: "zero duration fits",
			events: []event{
				{name: "E1", start: parseHM("09:00"), end: parseHM("09:59"), value: 60},
				{name: "Zero", start: parseHM("10:00"), end: parseHM("10:00"), value: 50},
				{name: "E2", start: parseHM("10:01"), end: parseHM("11:00"), value: 60},
			},
			want: []string{"E1", "Zero", "E2"},
		},
		{
			name: "negative value excluded",
			events: []event{
				{name: "A", start: parseHM("09:00"), end: parseHM("11:00"), value: 100},
				{name: "Neg", start: parseHM("11:00"), end: parseHM("12:00"), value: -100},
				{name: "B", start: parseHM("12:00"), end: parseHM("14:00"), value: 100},
			},
			want: []string{"A", "B"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := planTasks(tc.events)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func parseHM(s string) time.Time {
	layout := "15:04"
	t, err := time.Parse(layout, s)
	if err != nil {
		panic(err)
	}
	return t
}
