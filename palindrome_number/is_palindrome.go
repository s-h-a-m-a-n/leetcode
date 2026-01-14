package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	i := x
	j := 0
	for x != 0 {
		j = j*10 + x%10
		x /= 10
	}

	return i == j
}
