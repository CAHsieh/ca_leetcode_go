package main

func isValid(s string) bool {
	runes := []rune(s)

	st := []rune{}

	for _, c := range runes {
		if '(' == c || '[' == c || '{' == c {
			st = Push(c, st)
		} else if ')' == c {
			if len(st) > 0 && '(' == Peek(st) {
				st, _ = Pop(st)
			} else {
				return false
			}
		} else if ']' == c {
			if len(st) > 0 && '[' == Peek(st) {
				st, _ = Pop(st)
			} else {
				return false
			}
		} else if '}' == c {
			if len(st) > 0 && '{' == Peek(st) {
				st, _ = Pop(st)
			} else {
				return false
			}
		}
	}

	return (0 == len(st))
}

//Push add element
func Push(x rune, s []rune) []rune {
	s = append(s, x)
	return s
}

//Peek find top element
func Peek(s []rune) rune {
	return s[len(s)-1]
}

//Pop pop out top element
func Pop(s []rune) ([]rune, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}
