package main

func groupAnagrams(strs []string) [][]string {

	strsCode := make(map[string]int)
	ansMap := make(map[int][]string)

	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

	for _, str := range strs {
		key := 1
		for _, char := range str {
			key *= prime[char-'a']
		}
		strsCode[str] = key

		_, contain := ansMap[key]
		if !contain {
			ansMap[key] = make([]string, 0)
		}
		ansMap[key] = append(ansMap[key], str)
	}

	ans := make([][]string, 0)

	for k := range ansMap {
		ans = append(ans, ansMap[k])
	}

	return ans
}
