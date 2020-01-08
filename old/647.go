package main

func countSubstrings(s string) int {

	is := 0

	for i := 2; i <= len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			if isPalindromeString(s[j : j+i]) {
				is++
			}
		}
	}
	return (is + len(s))
}

func isPalindromeString(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
