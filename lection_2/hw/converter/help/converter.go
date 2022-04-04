package converter

type converter struct {
	symbols []string
	values  []int
}

type Converter interface {
	IntToRoman(num int) string
}

func (c *converter) IntToRoman(num int) string {
	//TODO реализовать конвертацию
}

func NewConverter(values []int, symbols []string) Converter {
	//TODO реализовать конструктор
}
