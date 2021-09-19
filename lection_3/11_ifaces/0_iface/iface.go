package main

import "fmt"

type Flyer interface {
	Fly()
}

type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name + " is flying")
}

func (b Bird) Greet() {
	fmt.Println("Hey there")
}

func DoFly(f Flyer) {

	f.Fly()
}

type Mig45 struct{}

func (m Mig45) Fly() {
	fmt.Println("Mig Flied away")
}

func main() {
	flSlice := []Flyer{
		Bird{"Duck plane"},
		Mig45{},
	}
	fmt.Println(flSlice)

	for _, v := range flSlice {
		v.Fly()
	}
	return
	duckPlane := Bird{"Duck plane"}

	GoFly(duckPlane)
}

func GoFly(f Flyer) {
	f.Fly()
	//b := f.(Bird)

	if b, ok := f.(Bird); ok {
		fmt.Println(b.Name)
	}
}
