package main

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	peopleMap := make(map[int][]int)

	for i := 0; i < len(people); i++ {
		mList, ok := peopleMap[people[i][0]]
		if !ok {
			mList = []int{people[i][1]}
		} else {
			mList = append(mList, people[i][1])
		}
		peopleMap[people[i][0]] = mList
	}

	var keys []int
	for k := range peopleMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	ans := make([][]int, 0)

	for _, k := range keys {
		mList := peopleMap[k]
		sort.Ints(mList)
		for _, p := range mList {
			ans = append(ans, []int{k, p})
			copy(ans[p+1:], ans[p:])
			ans[p] = []int{k, p}
		}
	}

	return ans
}
