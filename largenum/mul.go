package largenum

import "log"

func (a LargeNum) Mul(b LargeNum) LargeNum {
	positive := true
	if a.positive != b.positive {
		positive = false
	}

	a.positive = true
	b.positive = true

	result := a.mul(b)
	result.positive = positive
	return result
}

func (a LargeNum) mul(b LargeNum) LargeNum {
	if !a.positive && !b.positive {
		log.Panicf("Require positive multiplicands %s * %s", a, b)
	}

	indexA := len(a.fraction)
	indexB := len(b.fraction)
	a = a.Exp(indexA)
	b = b.Exp(indexB)

	results := make([]LargeNum, 0, 10)

	for i, cb := range b.integer {
		results = append(results, a.mulDigit(cb).Exp(i))
	}

	ln := Adds(results...)
	ln = ln.Exp(-(indexA + indexB))

	return ln
}

func (a LargeNum) mulDigit(b rune) LargeNum {
	if b < '0' || b > '9' {
		log.Panicf("Invalid multiplicand (not a digit) %c", b)
	}

	if len(a.fraction) != 0 || !a.positive {
		log.Panicf("Invalid multiplicand (not a positive integer) %s", a)
	}

	ln := initializeEmptyLargeNum()

	var memory rune = '0'
	var result rune

	for _, c := range a.integer {
		memory, result = digitmul(b, c, memory)
		ln.integer += string(result)
	}

	if memory > '0' {
		ln.integer += string(memory)
	}

	return ln
}
