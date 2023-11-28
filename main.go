package main

import (
	"fmt"
	"math/rand"
	"time"
)

var missionCompleted bool

func main() {
	checkMissionCompletion()
}

func checkMissionCompletion() {
	if missionCompleted {
		fmt.Println("Mission is now completed")
	} else {
		fmt.Println("Mission was failure.")
	}
}

func markMissionCompleted() {
	missionCompleted = true
}

func foundTreasure() bool {
	rand.NewSource(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
