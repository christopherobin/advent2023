package utils

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func ReadNumber(src []byte) (int, int) {
	number := 0
	n := 0
	for idx, c := range src {
		if !IsDigit(c) {
			break
		}

		n = idx
		number = number*10 + int(c-'0')
	}
	return number, n
}

func ParseNumbers(src []byte) []int {
	numbers := []int{}
	for i := 0; i < len(src); i++ {
		c := src[i]
		if !IsDigit(c) {
			continue
		}

		number, n := ReadNumber(src[i:])
		numbers = append(numbers, number)
		i += n
	}
	return numbers
}

func ParseNumbersString(src string) []int {
	return ParseNumbers([]byte(src))
}
