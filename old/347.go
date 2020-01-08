package main

func topKFrequent(nums []int, k int) []int {

	countMap := make(map[int]int)

	for _, n := range nums {
		val, ok := countMap[n]
		if ok {
			countMap[n] = val + 1
		} else {
			countMap[n] = 1
		}
	}

	s := make([]int, 0)

	for key := range countMap {
		insertIdx := len(s)
		for i, k := range s {
			if countMap[k] < countMap[key] {
				insertIdx = i
				break
			}
		}
		s = append(s, 0)
		copy(s[insertIdx+1:], s[insertIdx:])
		s[insertIdx] = key

		if len(s) > k {
			s = s[:len(s)-1]
		}
	}

	return s
}
