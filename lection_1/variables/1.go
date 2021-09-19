package main

import (
	"fmt"
	"strconv"
)

var (
	counter int
)

var (
	FirstName string = "dmitrii"
	SecName   string = "pashkov"
	Company   string = "mafin"
)

var i int = 52

func main() {
	var i int
	i = 42

	var f float32 = 32.4

	s1 := "Lets start with go"

	fmt.Printf("%v,%T \n", i, i)
	fmt.Printf("%v,%T \n", f, f)
	fmt.Printf("%v,%T \n", s1, s1)

	i = int(f)
	fmt.Printf("%v,%T \n", i, i)

	var s string
	s = strconv.Itoa(i)
	fmt.Printf("%v,%T \n", s, s)
}
