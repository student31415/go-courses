package main

import "fmt"

type Storage interface {
	Get(key int) (data int)
	Set(key int, data int)
}

type storageSlice struct {
	data []int
}

func (s *storageSlice) Get(key int) (data int) {
	return s.data[key]
}

func (s *storageSlice) Set(key int, data int) {
	s.data[key] = data
}

func NewStorageSlice() Storage {
	return &storageSlice{
		data: make([]int, 10),
	}
}

type storageMap struct {
	data map[int]int
}

func (s *storageMap) Get(key int) (data int) {
	return s.data[key]
}

func (s *storageMap) Set(key int, data int) {
	s.data[key] = data
}

func NewStorageMap() Storage {
	return &storageMap{
		data: make(map[int]int),
	}
}

func main() {
	//init storage
	storage1 := NewStorageSlice() //&{map[1:1 2:4]}

	//set value
	storage1.Set(1, 1)
	storage1.Set(2, 4)
	//...dsvsdvsdvsdv..//
	// get value

	fmt.Println(storage1.Get(1))
	fmt.Println(storage1.Get(2))
	//fmt.Println(storage1.Get(-1))
	fmt.Println(storage1)
}
