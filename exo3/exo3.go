package exo3

import "sort"

type ByEnd [][]int

func (a ByEnd) Len() int {
	return len(a)
}

func (a ByEnd) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByEnd) Less(i, j int) bool {
	return a[i][1] < a[j][1]
}

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Sort(ByEnd(intervals))

	nonOverlapCount := 1
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			nonOverlapCount++
			end = intervals[i][1]
		}
	}

	return len(intervals) - nonOverlapCount
}
