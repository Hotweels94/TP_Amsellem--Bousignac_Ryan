package main

import (
	"fmt"
	"sort"
)

func Ft_profit(prices []int) int {
	littleToBig := make([]int, len(prices))
	copy(littleToBig, prices)
	sort.Ints(littleToBig)
	lowestPrice := littleToBig[0]

	var startIndex int
	for i := 0; i < len(prices); i++ {
		if prices[i] == lowestPrice {
			startIndex = i
			break
		}
	}

	newArray := prices[startIndex:]
	sort.Sort(sort.Reverse(sort.IntSlice(newArray)))
	value := newArray[0] - lowestPrice

	fmt.Println(value)
	return value
}

func main() {
	Ft_profit([]int{7, 1, 5, 3, 6, 4}) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	Ft_profit([]int{7, 6, 4, 3, 1}) // resultat : 0
}
