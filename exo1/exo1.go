package exo1

import (
	"sort"
)

func Ft_coin(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	n := 0

	for _, coin := range coins {
		for amount >= coin {
			amount -= coin
			n++
		}
	}

	if amount != 0 {
		return -1
	}

	return n
}
