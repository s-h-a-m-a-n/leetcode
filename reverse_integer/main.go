package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	if x < -2147483648 {
		return 0
	}
	if x > 2147483647 {
		return 0
	}
	var s string
	sign := 1

	if x < 0 {
		sign = -1
		x = -x
	}
	s = strconv.Itoa(x)

	sb := bytes.Buffer{}
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}

	s = sb.String()
	x, _ = strconv.Atoi(s)

	x = sign * x

	if x < -2147483648 {
		return 0
	}
	if x > 2147483647 {
		return 0
	}

	return x
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(-120))
}
