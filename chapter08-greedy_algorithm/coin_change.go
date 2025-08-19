package chapter08_greedy_algorithm

import (
	"math"
)

type CurrencyUnit struct {
	name  string
	value int
}

type CurrencyUnits []CurrencyUnit

var currencyUnits = []CurrencyUnit{
	{name: "5грн.", value: 500},
	{name: "2грн.", value: 200},
	{name: "1грн.", value: 100},
	{name: "50к.", value: 50},
	{name: "25к.", value: 25},
	{name: "10к.", value: 10},
	{name: "5к.", value: 5},
	{name: "2к.", value: 2},
	{name: "1к.", value: 1},
}

func GetChange(sum float64) []string {
	if sum <= 0 {
		return []string{}
	}

	change := make([]string, 0)
	sumCoins := int(math.Round(sum * 100))

	for _, unit := range currencyUnits {
		count := sumCoins / unit.value
		if count == 0 {
			continue
		}
		for i := 0; i < count; i++ {
			change = append(change, unit.name)
		}
		sumCoins %= unit.value
	}

	return change
}
