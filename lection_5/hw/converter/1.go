package converter

type converter struct {
	symbols []string
	values  []int
}

type Converter interface {
	IntToRoman(num int) string
}

func (c *converter) IntToRoman(num int) string {
	res, i := "", 0
	for num != 0 {
		for c.values[i] > num {
			i++
		}
		num -= c.values[i]
		res += c.symbols[i]
	}
	return res
}

func NewConverter(values []int, symbols []string) Converter {
	return &converter{
		values:  values,
		symbols: symbols,
	}
}
