package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(reorderedPowerOf2(1))
	fmt.Println(reorderedPowerOf2(10))
	fmt.Println(reorderedPowerOf2(16))
	fmt.Println(reorderedPowerOf2(24))
	fmt.Println(reorderedPowerOf2(46))
}

func reorderedPowerOf2(N int) bool {
	table := []string{}
	for i := 1; i <= 1000_000_000; i *= 2 {
		str := strconv.Itoa(i)
		r := []rune(str)
		sort.Sort(sortRunes(r))
		table = append(table, string(r))
	}

	str := ""
	for N > 0 {
		str += strconv.Itoa(N % 10)
		N /= 10
	}

	r := []rune(str)
	sort.Sort(sortRunes(r))
	str = string(r)

	for _, t := range table {
		if str == t {
			return true
		}
	}

	return false
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
