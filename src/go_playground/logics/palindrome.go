package logics

func Palindrome(pal int) string {
	temp := pal
	var sum int = 0
	for pal > 0 {
		r := pal % 10
		sum = (sum * 10) + r
		pal /= 10
	}

	if temp == sum {
		return "it's palindrome number"
	}
	return "it's not palindrome number"

}
