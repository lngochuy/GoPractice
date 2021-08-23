package largenum

import (
	"log"
)

func (a LargeNum) Div(b LargeNum, length int) LargeNum {
	if b.Compare(Int2LargeNum(0)) == 0 {
		log.Panicln("Division by zero")
	}

	positive := true
	if a.positive != b.positive {
		positive = false
	}

	a.positive = true
	b.positive = true

	result := a.div(b, length)
	result.positive = positive
	return result
}

func (a LargeNum) div(b LargeNum, length int) LargeNum {
	if !a.positive || !b.positive {
		log.Panicf("Require positive numbers %s %s", a, b)
	}

	result := initializeEmptyLargeNum()
	indexA := len(a.fraction)
	indexB := len(b.fraction)
	a = a.Exp(indexA)
	b = b.Exp(indexB)

	positionA := a.ilen()
	tmpA := initializeEmptyLargeNum()

	for {
		if length >= positionA-(indexA-indexB) {
			break
		}

		positionA--
		if positionA >= 0 {
			tmpA.integer = string(a.iget(positionA)) + tmpA.integer
		} else {
			tmpA = tmpA.Exp(1)
		}

		if tmpA.Compare(b) == -1 {
			result.integer = "0" + result.integer
		} else {
			c := tmpA.divDigit(b)
			tmpA = tmpA.sub(b.Mul(c))
			result = result.appendInt(c)
		}
	}

	result = result.Exp(-(indexA - positionA - indexB))
	return result
}

func (a LargeNum) divDigit(b LargeNum) LargeNum {
	if !a.positive || !b.positive {
		log.Panicf("Require positive numbers %s %s", a, b)
	}

	if len(a.fraction) > 0 || len(b.fraction) > 0 {
		log.Panicf("Require integer numbers %s %s", a, b)
	}

	if a.Compare(b) == -1 {
		log.Panicf("Require %s >= %s", a, b)
	}

	i := 0
	var result LargeNum
	for {
		i++
		if i > 10 {
			log.Panicf("Invalid digit division, result is not a digit %d", i)
		}

		c := b.Mul(Int2LargeNum(i))

		comparison := c.Compare(a)
		if comparison == 0 {
			result = Int2LargeNum(i)
		}

		if comparison == 1 {
			result = Int2LargeNum(i - 1)
		}

		if comparison != -1 {
			break
		}
	}

	return result
}
