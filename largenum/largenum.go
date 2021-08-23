package largenum

import (
	"errors"
	"fmt"
	"log"
)

type LargeNum struct {
	integer  string
	fraction string
	positive bool
}

func initializeEmptyLargeNum() LargeNum {
	return LargeNum{
		integer:  "",
		fraction: "",
		positive: true,
	}
}

func InitializeLargeNum(num string) (LargeNum, error) {
	if len(num) == 0 {
		return initializeEmptyLargeNum(), nil
	}

	var i = 0
	var havedot = false
	var ln = LargeNum{fraction: "", integer: ""}
	ln.positive = !(num[0] == '-')

	if !ln.positive {
		i++
	}

	for _, c := range num[i:] {
		if c == '.' {
			if havedot {
				return ln, errors.New("invalid number: more than one dot character")
			}

			havedot = true
		} else if !isdigit(c) {
			return ln, errors.New("invalid number: non-digit character")
		} else {
			if !havedot {
				ln.integer = string(c) + ln.integer
			} else {
				ln.fraction += string(c)
			}
		}
	}

	ln.beautify()
	return ln, nil
}

func Int2LargeNum(a int) LargeNum {
	ln := initializeEmptyLargeNum()
	if a < 0 {
		a = -a
		ln.positive = false
	}

	for a > 0 {
		digit := rune(a % 10)
		a = a / 10
		ln.integer += string(digit + '0')
	}

	return ln
}

func (ln LargeNum) ilen() int {
	return len(ln.integer)
}

func (ln LargeNum) flen() int {
	return len(ln.fraction)
}

func (ln LargeNum) iget(index int) rune {
	if index < 0 || index >= len(ln.integer) {
		return '0'
	}

	return rune(ln.integer[index])
}

func (ln LargeNum) fget(index int) rune {
	if index < 0 || index >= len(ln.fraction) {
		return '0'
	}

	return rune(ln.fraction[index])
}

// The Exp method returns the ln * 10e(index)
//
// Example:
//
// ln.Exp(3) = ln * 10e3 = ln * 1000
//
// ln.Exp(-3) = ln / 1000
func (ln LargeNum) Exp(index int) LargeNum {
	newln := ln

	positiveIndex := index
	if index < 0 {
		positiveIndex = -index
	}

	for i := 0; i < positiveIndex; i++ {
		if index > 0 {
			fraction := ln.fget(i)
			if len(newln.fraction) > 0 {
				newln.fraction = newln.fraction[1:]
			}
			newln.integer = string(fraction) + newln.integer
		} else {
			integer := ln.iget(i)
			if len(newln.integer) > 0 {
				newln.integer = newln.integer[1:]
			}
			newln.fraction = string(integer) + newln.fraction
		}
	}

	return newln
}

// appendInt appends two large numbers together. Note that these numbers must be positive integer.
func (a LargeNum) appendInt(b LargeNum) LargeNum {
	if len(a.fraction) > 0 || len(b.fraction) > 0 {
		log.Panicf("Require integers %s %s", a, b)
	}

	if !a.positive || !b.positive {
		log.Panicf("Require positive numbers %s %s", a, b)
	}

	ln := a
	for i := b.ilen() - 1; i >= 0; i-- {
		ln.integer = string(b.integer[i]) + ln.integer
	}

	return ln
}

func (ln *LargeNum) beautify() {
	truncateZero := 0
	for i, c := range ln.integer {
		if c != '0' {
			truncateZero = i + 1
		}
	}

	ln.integer = ln.integer[:truncateZero]

	truncateZero = 0
	for i, c := range ln.fraction {
		if c != '0' {
			truncateZero = i + 1
		}
	}

	ln.fraction = ln.fraction[:truncateZero]

	if len(ln.integer) == 0 && len(ln.fraction) == 0 {
		ln.integer = "0"
	}
}

func (ln LargeNum) String() string {
	ln.beautify()

	integer := ""
	for _, c := range ln.integer {
		integer = string(c) + integer
	}

	fraction := ln.fraction

	if len(fraction) == 0 && (integer == "0" || len(integer) == 0) {
		return "0"
	}

	if !ln.positive {
		integer = "-" + integer
	}

	if len(fraction) != 0 {
		integer = integer + "." + ln.fraction
	}

	return integer
}

func (ln LargeNum) Debug() string {
	return fmt.Sprintf("%s %s %t", ln.integer, ln.fraction, ln.positive)
}

func (ln LargeNum) Compare(another LargeNum) int {
	if ln.positive {
		if !another.positive {
			return 1
		} else {
			return ln.compare(another)
		}
	} else {
		if another.positive {
			return -1
		} else {
			return another.compare(ln)
		}
	}
}

func (ln LargeNum) compare(another LargeNum) int {
	if ln.positive != another.positive {
		log.Panicln("Invalid comparison")
	}

	for i := max(ln.ilen(), another.ilen()) - 1; i >= 0; i-- {
		if r := compare(int(ln.iget(i)), int(another.iget(i))); r != 0 {
			return r
		}
	}

	for i := 0; i < max(ln.flen(), another.flen()); i++ {
		if r := compare(int(ln.fget(i)), int(another.fget(i))); r != 0 {
			return r
		}
	}

	return 0
}
