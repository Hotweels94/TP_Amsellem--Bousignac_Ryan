package exo5

import (
	"fmt"
)

func Ft_max_substring(s string) int {
	a := []rune(s)
	b := 0
	for i := 0; i < len(s); i++ {
		b = b + 1
		if a[0] == a[b] {
			break
		} else {
			b = b + 1
			continue
		}
	}
	fmt.Println(b)
	return b
}
