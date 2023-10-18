package main

import (
	"fmt"
)

func sqrter(x int) int {
	var ab int
	for i := 0; i < x; i++ {
		if x-(i*i) >= 0 {
			ab = i
		} else if x-(i*i) < 0 {
			break
		}
	}
	return ab
}

func main() {
	fmt.Println(sqrter(123798))
}
