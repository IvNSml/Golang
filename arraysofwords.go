package main

import (
	"fmt"
)

func main() {
	var s1 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	var s2 = []string{"bbbaaayddqbbrrrv"}
	fmt.Println(MxDifLg(s1, s2))
}
func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	var (
		min1 = len(a1[0])
		max1 = len(a1[0])
		min2 = len(a2[0])
		max2 = len(a2[0])
	)

	for _, value := range a1 {
		if len(value) > max1 {
			max1 = len(value)
		}
		if len(value) < min1 {
			min1 = len(value)
		}
	}
	for _, value := range a2 {
		if len(value) > max2 {
			max2 = len(value)
		}
		if len(value) < min2 {
			min2 = len(value)
		}
	}
	if max1-min2 > max2-min1 {
		return max1 - min2
	} else {
		return max2 - min1
	}
}
