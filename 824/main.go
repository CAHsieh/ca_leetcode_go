package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
	fmt.Println(toGoatLatin("The quick brown fox jumped over the lazy dog"))
}

func toGoatLatin(S string) string {
	arr := strings.Split(S, " ")
	a := "a"
	for i := 0; i < len(arr); i++ {

		if arr[i][0] == 'a' || arr[i][0] == 'e' || arr[i][0] == 'i' || arr[i][0] == 'o' || arr[i][0] == 'u' ||
			arr[i][0] == 'A' || arr[i][0] == 'E' || arr[i][0] == 'I' || arr[i][0] == 'O' || arr[i][0] == 'U' {
			arr[i] += "ma"
		} else {
			arr[i] = string(arr[i][1:]) + string(arr[i][0]) + "ma"
		}

		arr[i] += a
		a += "a"
	}
	return strings.Join(arr, " ")
}
