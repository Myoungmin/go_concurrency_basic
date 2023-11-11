package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	go throwingNinjaStar(channel)
	for {
		message, open := <-channel
		if !open {
			fmt.Println("channel is closed!")
			break
		}
		fmt.Println(message)
	}
}

func throwingNinjaStar(channel chan string) {
	rand.NewSource(time.Now().UnixNano())
	numRounds := 3
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}
	// 아래의 Close를 하지 않으면 deadLock 발생
	close(channel)
}
