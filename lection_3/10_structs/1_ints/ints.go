package main

import (
	"fmt"
	"sort"
)

type MyInt int

func (m MyInt) showYourSelf() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	*m = *m + MyInt(i)
}

type MyStruct struct {
	Num  int
	Name string
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
	i.showYourSelf()
	i.add(3)
	i.showYourSelf()
	a := mySliceStruct{
		{
			Num:  4,
			Name: "MyStruct4",
		},
		{
			Num:  -1,
			Name: "MyStruct2",
		},
		{
			Num:  1,
			Name: "MyStruct1",
		},
		{
			Num:  3,
			Name: "MyStruct3",
		},
	}
	fmt.Println(a)
	sorter(a)
	fmt.Println(a)
	a.Swap(0, 3)
	fmt.Println(a)
	a.Swap(0, 3)
	fmt.Println(a)
	fmt.Println("1:", a.Less(1))
	fmt.Println("0:", a.Less(0))
	fmt.Println(a.Len())
}

func sorter(a []MyStruct) []MyStruct {
	sort.Slice(a, func(i, j int) bool {
		return a[i].Num < a[j].Num
	})
	return a
}
