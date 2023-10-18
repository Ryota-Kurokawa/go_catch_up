package main

import (
	"fmt"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func sumroop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// main関数

func main() {
	fmt.Println("My favorite number is", rand.Intn(300))
	fmt.Println(add(42, 13))
	a, b := swap(1, 2)
	fmt.Println(a, b)
	sumroop()
}
