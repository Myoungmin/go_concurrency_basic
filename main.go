package main

import (
	"sync"
)

func main() {
	// var beeper sync.WaitGroup
	// evilNinjas := []string{"Tommy", "Johnny", "Bobby"}
	// beeper.Add(len(evilNinjas))

	// for _, evilNinja := range evilNinjas {
	// 	go attack(evilNinja, &beeper)
	// }

	// beeper.Wait()
	// fmt.Println("Mission completed")

	var beeper sync.WaitGroup
	beeper.Add(1)
	// 1개만 Done 되어야한다.
	// 없으면 deadLock, 2개 이상되면 negative WaitGroup couter
	beeper.Done()
	//beeper.Done()
	beeper.Wait()
}

// func attack(evilNinja string, beeper *sync.WaitGroup) {
// 	fmt.Println("Attacked evil ninja:", evilNinja)
// 	beeper.Done()
// }
