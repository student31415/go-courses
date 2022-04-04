package main

import "fmt"

type person struct {
	Name string
	inn  string
}

type Stuff struct {
	inn int
}

type SecretAgent struct {
	person
	Stuff
	LicenseToKill bool
}

func (p person) GetName() string {
	return p.Name
}

func (s SecretAgent) GetName() string {
	return "CLASSIFIED"
}

func main() {
	sa := SecretAgent{
		person: person{
			"James",
			"12312321321",
		},
		LicenseToKill: true,
	}

	fmt.Printf("%T %+v\n", sa, sa)
	fmt.Println("secret inn", sa.GetName())
	fmt.Println("secret inn", sa.person.GetName())
	fmt.Println(sa.Name)

}
