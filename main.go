package main

import (
	"fmt"
)

func main() {
	go attack("Tommy")
	fmt.Println("Mission completed")
}

func attack(evilNinja string) {
	fmt.Println("Attacked evil ninja:", evilNinja)
}
