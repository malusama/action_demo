package main

import (
	"log"
)

func add(x, y int) int {
	return x + y
}

func main() {
	log.Println("hello world")

	ret_xx := []int{1, 2, 3, 4}
	for v := range []int{1, 2, 3} {
		ret_xx = append(ret_xx, v)
	}

	log.Printf("%v", ret_xx)
}
