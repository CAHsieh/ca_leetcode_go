package main

func generateParenthesis(n int) []string {
	ans := make([]string, 0)

	leftStr := ""

	for i := 0; i < n; i++ {
		leftStr = leftStr + "("
	}

	generate(1, leftStr, n, &ans)

	return ans
}

func generate(at int, str string, n int, ans *[]string) {
	if at > len(str) {
		if len(str) == n*2 {
			*ans = append(*ans, str)
		}
		return
	}

	count := 0

	for i := 0; i < at && i < len(str); i++ {
		if string(str[i]) == "(" {
			count++
		} else {
			count--
		}
	}

	for i := 0; i <= count; i++ {
		tmp := str
		for j := 0; j < i; j++ {
			tmp = tmp[:at] + ")" + tmp[at:]
		}
		generate(at+i+1, tmp, n, ans)
	}
}
