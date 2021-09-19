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
}

func getTimer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time from start %v", time.Since(start))
	}
}
