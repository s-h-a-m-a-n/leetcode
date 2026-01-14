package string_to_integer_atoi

import "math"

func myAtoi(s string) int {
	if len(s) == 0 || len(s) > 200 {
		return 0
	}
	sign := 1
	runes := []rune(s)
	for 0 < len(runes) {
		r := runes[0]
		switch {
		case r == '+':
			runes = runes[1:]
			goto label
		case r == '.':
			return 0
		case r >= 'a' && r <= 'z':
			return 0
		case r >= 'A' && r <= 'Z':
			return 0
		case r >= '0' && r <= '9':
			goto label
		case r == ' ':
			runes = runes[1:]
		case r == '-':
			sign = -1
			runes = runes[1:]
			goto label
		}
	}
label:
	res := 0
	for 0 < len(runes) {
		r := runes[0]
		runes = runes[1:]
		switch r {
		case '0':
			res = res * 10
		case '1':
			res = res*10 + 1
		case '2':
			res = res*10 + 2
		case '3':
			res = res*10 + 3
		case '4':
			res = res*10 + 4
		case '5':
			res = res*10 + 5
		case '6':
			res = res*10 + 6
		case '7':
			res = res*10 + 7
		case '8':
			res = res*10 + 8
		case '9':
			res = res*10 + 9
		default:
			goto label2
		}
		if res > math.MaxInt32 {
			res = math.MaxInt32 + 1
		}
	}
label2:

	res = res * sign

	if res > math.MaxInt32 {
		return math.MaxInt32
	} else if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}
