package main

import "fmt"

func main() {
	var b1 bool
	n := 1 == 1
	fmt.Printf("%v,%T \n", n, n)
	fmt.Printf("%v,%T \n", b1, b1)

	var i int = 32
	fmt.Printf("%v,%T \n", i, i)

	a := 10
	b := 3
	fmt.Println(a + a)
	fmt.Println(a - a)
	fmt.Println(a * a)
	fmt.Println(a / b)
	fmt.Println(a % b)
}
