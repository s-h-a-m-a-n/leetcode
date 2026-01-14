package main

import (
	"fmt"
)

func isValid(s string) bool {
	dict := map[rune]int8{
		')': 0,
		']': 1,
		'}': 2,
	}

	var r []int8
	for _, v := range s {
		switch v {
		case '(':
			r = append(r, 0)
		case '[':
			r = append(r, 1)
		case '{':
			r = append(r, 2)
		default:
			if len(r) == 0 || r[len(r)-1] != dict[v] {
				return false
			}
			r = r[:len(r)-1]
		}
	}
	if len(r) == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isValid("(("))
	fmt.Println(isValid("("))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("((()))"))
	fmt.Println(isValid("((()])"))
}
