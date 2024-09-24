package exo2

import (
	"fmt"
	"sort"
)

func Ft_missing(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i {
			fmt.Println(i)
			return i
		} else if nums[n-1] == i {
			fmt.Println(i + 1)
			return i + 1
		}
	}
	return n
}
