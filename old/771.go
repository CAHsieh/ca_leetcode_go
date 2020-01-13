package main

func numJewelsInStones(J string, S string) int {
	count := make([]int, 127)
	sum := 0

	for i := 0; i < len(J); i++ {
		count[J[i]]++
	}

	for i := 0; i < len(S); i++ {
		sum += count[S[i]]
	}

	return sum
}
