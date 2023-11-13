package main

import "fmt"

func main() {
	ninja1, ninja2 := make(chan string), make(chan string)

	go captainElect(ninja1, "Ninja 1")
	go captainElect(ninja2, "Ninja 2")

	fmt.Println(<-ninja1)
	fmt.Println(<-ninja2)
}

func captainElect(ninja chan string, message string) {
	ninja <- message
}
