package main

import "fmt"

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
	fmt.Println(maxNumberOfBalloons("leetcode"))
	fmt.Println(maxNumberOfBalloons("lloo"))
}

func maxNumberOfBalloons(text string) int {
	m := make(map[byte]int)
	m['b'] = 0
	m['a'] = 0
	m['l'] = 0
	m['o'] = 0
	m['n'] = 0
	max := len(text)
	for i := 0; i < max; i++ {
		if text[i] == 'b' || text[i] == 'a' || text[i] == 'l' || text[i] == 'o' || text[i] == 'n' {
			v, _ := m[text[i]]
			m[text[i]] = v + 1
		}
	}

	for k, v := range m {
		count := v
		if k == 'l' || k == 'o' {
			count = v / 2
		}
		if count < max {
			max = count
		}
	}
	return max
}
