package largenum

import "log"

func Adds(lns ...LargeNum) LargeNum {
	if len(lns) == 0 {
		return initializeEmptyLargeNum()
	}

	sum := lns[0]
	for _, ln := range lns[1:] {
		sum = sum.Add(ln)
	}

	return sum
}

func (a LargeNum) Add(b LargeNum) LargeNum {
	if !a.positive {
		if b.positive {
			// (-a) + (+b) = b - (+a)
			a.positive = true
			return b.sub(a)
		} else {
			// (-a) + (-b) = -[(+a) + (+b)]
			a.positive = true
			b.positive = true
			result := a.add(b)
			result.positive = false
			return result
		}
	} else {
		if b.positive {
			// (+a) + (+b) (normally)
			return a.add(b)
		} else {
			// (+a) + (-b) = (+a) - (+b)
			b.positive = true
			return a.sub(b)
		}
	}
}

func (a LargeNum) add(b LargeNum) LargeNum {
	if !a.positive || !b.positive {
		log.Panicf("add() method only works on positive numbers")
	}

	var sum = initializeEmptyLargeNum()
	var memory rune = '0'
	var result rune

	// fraction part
	n := max(a.flen(), b.flen())
	for i := n - 1; i >= 0; i-- {
		f1 := a.fget(i)
		f2 := b.fget(i)

		memory, result = digitsum(f1, f2, memory)

		sum.fraction = string(result) + sum.fraction
	}

	// integer part
	n = max(a.ilen(), b.ilen())
	for i := 0; i < n; i++ {
		i1 := a.iget(i)
		i2 := b.iget(i)

		memory, result = digitsum(i1, i2, memory)

		sum.integer += string(result)
	}

	if memory > '0' {
		sum.integer += string(memory)
	}

	return sum
}
