package converter

import (
	"strings"
)

func intToRoman(num int, one, five, ten rune) (res string) {
	// spliting num = a1a2a4...an-1an into  a1a2a4...an-1 and an
	first, last := num/10, num%10

	// transforming	last to roman
	switch {
	case last == 0:
		res = ""
	case last == 4:
		res = string(one) + string(five)
	case last == 5:
		res = string(five)
	case last == 9:
		res = string(one) + string(ten)

	case last < 4:
		res = strings.Repeat(string(one), last)
	default:
		res = string(five) + strings.Repeat(string(one), last-5)
	}
	// if first != 0 it means, that we are transforming 1000+ number from IntToRoman
	res = strings.Repeat(string(ten), first) + res

	return
}

func IntToRoman(dec int) (roman string) {
	ones, tens, hundrs := dec%10, (dec/10)%10, dec/100
	return intToRoman(hundrs, 'C', 'D', 'M') +
		intToRoman(tens, 'X', 'L', 'C') +
		intToRoman(ones, 'I', 'V', 'X')
}

type converter struct {
	symbols []string
	values  []int
}

type Converter interface {
	IntToRoman(num int) string
}

func (c *converter) IntToRoman(num int) (res string) {
	var div int
	res = ""
	for i := 0; num > 0; i++ {
		div = num / c.values[i]
		num = num % c.values[i]
		for div > 0 {
			div -= 1
			res += c.symbols[i]
		}
	}

	return
}

func NewConverter(values []int, symbols []string) Converter {
	return &converter{symbols, values}
}
