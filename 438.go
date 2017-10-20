package main

func findAnagrams(s string, p string) []int {

	pRunes := []rune(p)
	asciiCount := make([]int, 123)
	for _, char := range pRunes {
		asciiCount[char]++
	}

	var ans []int
	sRunes := []rune(s)
	left := 0
	right := 0
	count := len(p)
	for right < len(s) {
		if asciiCount[sRunes[right]] > 0 {
			count--
		}
		asciiCount[sRunes[right]]--
		right++

		if 0 == count {
			ans = append(ans, left)
		}

		if right-left == len(p) {
			if asciiCount[sRunes[left]] >= 0 {
				count++
			}
			asciiCount[sRunes[left]]++
			left++
		}
	}

	return ans
}
