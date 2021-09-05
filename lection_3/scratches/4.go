package main

import "fmt"

func main() {
	//nums:= []int{11,15,42,3,6,8,0,-1,2,7}
	//target:= 9
	//fmt.Println(twoSum(nums, target))
	//fmt.Println(max(nums))

	var a = 1
	var pa *int
	pa = &a
	fmt.Println(a)
	add(pa, 7)

	fmt.Println(a)
}

func add(a *int, v int) {
	*a = *a + v
}

func max(a []int) (index, value int) {
	maxV := -100
	maxIndex := 0
	for i := 0; i < len(a); i++ {
		if a[i] > maxV {
			maxV = a[i]
			maxIndex = i
		}
	}
	return maxV, maxIndex
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
