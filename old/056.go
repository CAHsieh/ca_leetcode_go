package main

import "sort"

func merge(intervals []Interval) []Interval {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var ans []Interval
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		check := intervals[i]
		if current.End >= check.Start {
			current.End = max(current.End, check.End)
		} else {
			ans = append(ans, current)
			current = check
		}
	}
	ans = append(ans, current)
	return ans
}

// func merge(intervals []Interval) []Interval {
// 	slice := make([]int, 0)
// 	slice2 := make([]rune, 0)
// 	for _, interval := range intervals {
// 		slice = append(slice, interval.Start)
// 		slice = append(slice, interval.End)
// 		slice2 = append(slice2, '(')
// 		slice2 = append(slice2, ')')
// 	}
// 	count := len(slice)
// 	for i := count - 1; i > 0; i-- {
// 		for j := 0; j < i; j++ {
// 			if slice[j] > slice[j+1] {
// 				tmp := slice[j]
// 				slice[j] = slice[j+1]
// 				slice[j+1] = tmp
// 				tmp2 := slice2[j]
// 				slice2[j] = slice2[j+1]
// 				slice2[j+1] = tmp2
// 			}
// 		}
// 	}
// 	for i := 0; i < count-1; i++ {
// 		if slice[i] == slice[i+1] && slice2[i] != slice2[i+1] {
// 			slice2[i] = '('
// 			slice2[i+1] = ')'
// 		}
// 	}
// 	trashmode := false
// 	trashcount := 0
// 	intervals = make([]Interval, 0)
// 	c := 0
// 	for i := count - 1; i >= 0; i-- {
// 		if !trashmode {
// 			c = slice[i]
// 			trashmode = true
// 		} else {
// 			if slice2[i] == '(' {
// 				if trashcount != 0 {
// 					trashcount--
// 				} else {
// 					intervals = append(intervals, Interval{slice[i], c})
// 					trashmode = false
// 				}
// 			} else {
// 				trashcount++
// 			}
// 		}
// 	}
// 	return intervals
// }
