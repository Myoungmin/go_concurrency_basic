package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	numRounds := 3
	go throwingNinjaStar(channel, numRounds)
	for i := 0; i < numRounds; i++ {
		fmt.Println(<-channel)
	}
}

func throwingNinjaStar(channel chan string, numRounds int) {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}
}
