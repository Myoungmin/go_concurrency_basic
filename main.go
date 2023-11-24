package main

import (
	"fmt"
	"time"
)

var (
	count int
)

func main() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Resulted count is: ", count)
}

func increment() {
	count++
}
