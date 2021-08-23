package largenum

import "log"

func isdigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}

func rune2digit(c rune) uint {
	if !isdigit(c) {
		log.Panicf("Invalid digit %c", c)
	}

	return uint(c - '0')
}

func digit2rune(d uint) rune {
	if d > 9 {
		log.Panicf("invalid digit %d", d)
	}

	return rune(d) + '0'
}

func max(nums ...int) int {
	m := nums[0]
	for _, n := range nums {
		if m < n {
			m = n
		}
	}

	return m
}

func digitsum(a, b, m rune) (memory rune, result rune) {
	if a > '9' || a < '0' {
		log.Panicf("Invalid digit %c", a)
	}

	if b > '9' || b < '0' {
		log.Panicf("Invalid digit %c", b)
	}

	if m > '9' || m < '0' {
		log.Panicf("Invalid memory %c", m)
	}

	// a + b + m = (adigit + '0') + (bdigit + '0') + (mdigit + '0')
	// => sum = a + b + m - 3 * '0'
	var sum = a + b + m - 3*'0'

	return sum/10 + '0', sum%10 + '0'
}

func digitsub(a, b, m rune) (memory rune, result rune) {
	if a > '9' || a < '0' {
		log.Panicf("Invalid digit %c", a)
	}

	if b > '9' || b < '0' {
		log.Panicf("Invalid digit %c", b)
	}

	if m > '9' || m < '0' {
		log.Panicf("Invalid memory %c", m)
	}

	// b + m = (bdigit + '0') + (mdigit + '0')
	// => b + m - '0'
	b = b + m - '0'

	memory = '0'
	if a < b {
		a += 10
		memory = '1'
	}

	c := a - b + '0'

	if c < '0' || c > '9' {
		log.Panicf("Invalid subtract result %c", c)
	}

	return memory, c
}

func digitmul(a, b, m rune) (memory rune, result rune) {
	if a > '9' || a < '0' {
		log.Panicf("Invalid digit multiplicand %c", b)
	}

	if b > '9' || b < '0' {
		log.Panicf("Invalid digit multiplicand %c", b)
	}

	if m > '9' || m < '0' {
		log.Panicf("Invalid digit memory %c", m)
	}

	c := (a-'0')*(b-'0') + (m - '0')

	return rune(c/10) + '0', rune(c%10) + '0'
}

func compare(a, b int) int {
	if a > b {
		return 1
	}

	if a < b {
		return -1
	}

	return 0
}
