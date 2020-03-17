package main

import "fmt"

func main() {
	s := []byte{'A', ' ', 'm', 'a', 'n', ',', ' ', 'a', ' ', 'p', 'l', 'a', 'n', ',', ' ', 'a', ' ', 'c', 'a', 'n', 'a', 'l', ':', ' ', 'P', 'a', 'n', 'a', 'm', 'a'}
	fmt.Println(s)
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

// func reverseString(s []byte) {
// 	count := len(s) - 1
// 	for i := 0; i < count; i++ {
// 		s[i], s[count] = s[count], s[i]
// 		count--
// 	}
// }
