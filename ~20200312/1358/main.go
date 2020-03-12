package main

import "fmt"

func main() {
	fmt.Println(numberOfSubstrings("abcabc"))
	fmt.Println(numberOfSubstrings("aaacb"))
	fmt.Println(numberOfSubstrings("abc"))
}

func numberOfSubstrings(s string) int {
	ans := 0
	i := 0
	t := -1
	count := make([]int, 3)
	length := len(s)
	for t < length {
		for count[0]*count[1]*count[2] == 0 {
			if t == length-1 {
				return ans
			}
			t++
			count[s[t]-'a']++
		}

		ans += length - t
		count[s[i]-'a']--
		i++
	}
	return ans
}
