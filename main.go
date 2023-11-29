package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionCompleted bool

func main() {
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			if foundTreasure() {
				markMissionCompleted()
			}
			wg.Done()
		}()
	}

	wg.Wait()

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
	fmt.Println("Marking mission completed.")
}

func foundTreasure() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
