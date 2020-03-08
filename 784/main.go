package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
	fmt.Println(letterCasePermutation("12345"))
}

func letterCasePermutation(S string) []string {
	if len(S) == 1 {
		if S[0] >= '0' && S[0] <= '9' {
			return []string{S}
		}
		return []string{strings.ToLower(S), strings.ToUpper(S)}
	}
	nS := S[1:]
	subRes := letterCasePermutation(nS)
	count := len(subRes)
	if S[0] >= '0' && S[0] <= '9' {
		result := make([]string, count)
		for i := 0; i < count; i++ {
			result[i] = string(S[0]) + subRes[i]
		}
		return result
	}

	result := make([]string, count*2)
	for i := 0; i < count*2; i += 2 {
		result[i] = strings.ToLower(string(S[0])) + subRes[i/2]
		result[i+1] = strings.ToUpper(string(S[0])) + subRes[i/2]
	}
	return result
}
