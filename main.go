package main

import (
	"log"
)

func add(x, y int) int {
	return x + y
}

func main() {
	log.Println("hello world")

	a := []int{1, 2, 3, 4}
	for v := range []int{1, 2, 3} {
		a = append(a, v)
	}

	log.Printf("%v", a)
}
