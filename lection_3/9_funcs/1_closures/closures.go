package main

import (
	"fmt"
	"time"
)

func main() {
	myTimer := getTimer()
	//myTimer()
	f := func() {
		myTimer()
	}

	f()
	f()
}

func getTimer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time from start %v\n", time.Since(start))
	}
}
