package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	go throwingNinjaStar(channel)
	fmt.Println(<-channel)
}

func throwingNinjaStar(channel chan string) {
	rand.NewSource(time.Now().UnixNano())
	score := rand.Intn(10)
	channel <- fmt.Sprint("You scored: ", score)
}
