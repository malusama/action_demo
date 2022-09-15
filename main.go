package main

import (
	"log"
)

func Add(x, y int) int {
	return x + y
}

func main() {
	log.Println("hello world")
	existing := make([]int64, 10, 10)

	var init []int64
	for _, element := range existing {
		init = append(init, element)
	}
	log.Printf("%v", init)
}
