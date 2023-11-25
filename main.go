package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock   sync.Mutex
	reLock sync.RWMutex
	count  int
)

func main() {
	//basics()
	readAndWrite()
}

func basics() {
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Resulted count is: ", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}

func readAndWrite() {
	go read()
	go write()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func read() {
	reLock.RLock()
	defer reLock.RUnlock()

	fmt.Println("Read locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Read unlocking")
}

func write() {
	reLock.Lock()
	defer reLock.Unlock()

	fmt.Println("Write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Writing unlocking")
}
