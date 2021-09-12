package main

import (
	"fmt"
	"sort"
)

type MyStruct struct {
	Num  int
	Name string
}

type MyInt int

func (m MyInt) showYourSelf() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	*m = *m + MyInt(i)
}

type mySliceStruct []MyStruct

func (m mySliceStruct) Less(i int) bool {
	return m[i].Num > 0
}

func (m mySliceStruct) Len() int {
	return len(m)
}
func (m mySliceStruct) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	i := MyInt(0)

	i.add(3)
	i.showYourSelf()
}

func sorter(a []MyStruct) []MyStruct {
	sort.Slice(a, func(i, j int) bool {
		return a[i].Num < a[j].Num
	})
	return a
}
