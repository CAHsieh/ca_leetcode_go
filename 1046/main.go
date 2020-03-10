package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight([]int{1, 3}))
	fmt.Println(lastStoneWeight([]int{3, 7, 8}))
}

func lastStoneWeight(stones []int) int {

	sort.Ints(stones)

	length := len(stones)

	for length > 1 {
		stones[length-1] -= stones[length-2]
		if stones[length-1] == 0 {
			stones = stones[:length-2]
			length -= 2
		} else if length > 2 {
			last := stones[length-1:]
			stones = stones[:length-2]
			for i := 0; i < length-2; i++ {
				if stones[i] >= last[0] {
					stones = append(stones[:i], append(last, stones[i:]...)...)
					break
				} else if i == length-3 {
					stones = append(stones, last[0])
					break
				}
			}
			length--
		} else {
			stones = []int{stones[1]}
			length = 1
		}
	}

	if len(stones) == 1 {
		return stones[0]
	}

	return 0
}
