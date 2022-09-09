package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("hello world")

	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}

	fmt.Printf("%v", arr)
}
