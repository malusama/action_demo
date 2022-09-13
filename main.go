package main

import (
	"log"
)

func Add(x, y int) int {
	return x + y
}

func main() {
	log.Println("hello world")

	a := []int{}
	for v := range []int{1, 2, 3} {
		a = append(a, v)
	}

	log.Printf("%v", a)
}
