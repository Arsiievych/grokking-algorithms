package chapter08_greedy_algorithm

import (
	"reflect"
	"testing"
)

func TestGetChange(t *testing.T) {
	type testCase struct {
		name string
		sum  float64
		want []string
	}

	cases := []testCase{
		{
			name: "zero_amount",
			sum:  0,
			want: []string{},
		},
		{
			name: "negative_amount",
			sum:  -5.50,
			want: []string{},
		},
		{
			name: "small_amount_nine_kopecks",
			sum:  0.09,
			want: []string{"5к.", "2к.", "2к."},
		},
		{
			name: "single_kopeck",
			sum:  0.01,
			want: []string{"1к."},
		},
		{
			name: "exact_kopeck_values",
			sum:  0.37,
			want: []string{"25к.", "10к.", "2к."},
		},
		{
			name: "mixed_hryvnia_and_kopecks",
			sum:  3.87,
			want: []string{"2грн.", "1грн.", "50к.", "25к.", "10к.", "2к."},
		},
		{
			name: "exact_hryvnia_amount",
			sum:  8.00,
			want: []string{"5грн.", "2грн.", "1грн."},
		},
		{
			name: "large_amount_with_change",
			sum:  13.42,
			want: []string{"5грн.", "5грн.", "2грн.", "1грн.", "25к.", "10к.", "5к.", "2к."},
		},
		{
			name: "only_largest_denomination",
			sum:  10.00,
			want: []string{"5грн.", "5грн."},
		},
		{
			name: "requires_all_denominations",
			sum:  8.88,
			want: []string{"5грн.", "2грн.", "1грн.", "50к.", "25к.", "10к.", "2к.", "1к."},
		},
		{
			name: "floating_point_precision_issue",
			sum:  1.99,
			want: []string{"1грн.", "50к.", "25к.", "10к.", "10к.", "2к.", "2к."},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetChange(tc.sum)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetChange() = %v, want %v", got, tc.want)
			}
		})
	}

}
