package main

func hammingDistance(x int, y int) int {

	n := 0
	for x > 0 || y > 0 {
		if x%2 != y%2 {
			n++
		}
		x /= 2
		y /= 2
	}

	return n
}
