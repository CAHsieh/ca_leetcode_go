package main

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}

	var digitMap = []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	ans := make([]string, 1)
	for i, s := range digits {
		for len(ans[0]) == i {
			str := ans[0]
			ans = ans[1:]

			for _, c := range digitMap[s-'0'] {
				ans = append(ans, str+string(c))
			}
		}
	}

	return ans
}
