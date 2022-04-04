package main

import "github.com/CodingSquire/go-courses/lection_3/8_package_func_doc/world"

func main() {

	//go get golang.org/x/tools/cmd/godoc
	//godoc -http=:6060
	world.PrintStartRoom()

	// обращение к переменной или константе пакета
	println("starting room: ", world.StartingRoom)

	// приватная переменная (с маленькой буквы) доступна только внутри пакета
	//println("starting level: ", world.startingLevel)
	// cannot refer to unexported name world.startingLevel

	println("starting level: ", world.GetStartingLevel())

	// к приватной функции обратиться нельзя
	println("check statring level: ", world.CheckStartingLevel(10))
	// cannot refer to unexported name world.checkStartingLevel

	// приватная область видимости распространяется на весь файл
	println("max level: ", maxLevel)

}
