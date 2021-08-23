package largenum

import "log"

func (a LargeNum) Sub(b LargeNum) LargeNum {
	if !a.positive {
		if b.positive {
			// (-a) - (+b) = - [(+a) + (+b))
			a.positive = true
			result := a.add(b)
			result.positive = false
			return result
		} else {
			// (-a) - (-b) = (-a) + (+b) = (+b) - (+a)
			a.positive = true
			b.positive = true
			return b.sub(a)
		}
	} else {
		if !b.positive {
			// (+a) - (-b) = a + (+b)
			b.positive = true
			return a.add(b)
		} else {
			// (+a) - (+b) (normally)
			return a.sub(b)
		}
	}
}

func (a LargeNum) sub(b LargeNum) LargeNum {
	if !a.positive || !b.positive {
		log.Panicf("sub() method only works on positive numbers")
	}

	if a.Compare(b) == -1 {
		result := b.Sub(a)
		result.positive = false
		return result
	}

	var diff = initializeEmptyLargeNum()
	var memory rune = '0'
	var result rune

	// fraction part
	n := max(a.flen(), b.flen())
	for i := n - 1; i >= 0; i-- {
		f1 := a.fget(i)
		f2 := b.fget(i)

		memory, result = digitsub(f1, f2, memory)

		diff.fraction = string(result) + diff.fraction
	}

	// integer part
	n = max(a.ilen(), b.ilen())
	for i := 0; i < n; i++ {
		i1 := a.iget(i)
		i2 := b.iget(i)

		memory, result = digitsub(i1, i2, memory)

		diff.integer += string(result)
	}

	if memory > '0' {
		// '10' - memory = (10 + '0') - (memorydigit + '0') + '0'
		// = 10 - memory + 2 * '0'
		diff.integer += string(10 - memory + '0' + '0')
		diff.positive = false
	}

	return diff
}
