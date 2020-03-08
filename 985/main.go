package main

import "fmt"

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}))
	fmt.Println(sumEvenAfterQueries([]int{5, 1}, [][]int{{0, 1}, {4, 0}}))
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	ans := []int{}
	for i, query := range queries {
		before := A[query[1]]
		A[query[1]] += query[0]
		if i == 0 {
			sum := 0
			for _, v := range A {
				if v%2 == 0 {
					sum += v
				}
			}
			ans = append(ans, sum)
		} else {
			if before%2 == 0 && A[query[1]]%2 != 0 {
				ans = append(ans, ans[i-1]-before)
			} else if before%2 == 0 && A[query[1]]%2 == 0 {
				ans = append(ans, ans[i-1]+query[0])
			} else if before%2 != 0 && A[query[1]]%2 == 0 {
				ans = append(ans, ans[i-1]+A[query[1]])
			} else {
				ans = append(ans, ans[i-1])
			}
		}
	}
	return ans
}
