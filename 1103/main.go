package main

import "fmt"

func main() {
	fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies(10, 3))
}

func distributeCandies(candies int, num_people int) []int {

	row := make([]int, num_people)

	i := 0
	n := 1
	for candies > 0 {
		if candies >= n {
			row[i] += n
		} else {
			row[i] += candies
		}

		candies -= n
		n++
		i++
		i %= num_people
	}

	return row
}
