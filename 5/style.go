package main

import "fmt"

type Vertex struct {
	x int
	y int
}

var A [2]Vertex

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	fmt.Println("---------------------------------------------------------------------")
	v := Vertex{1, 2}
	fmt.Println(v.x)
	fmt.Println(v.y)
	fmt.Println("---------------------------------------------------------------------")
	A[0] = Vertex{1, 2}
	A[1] = Vertex{3, 4}
	fmt.Println(A[0], A[1])

}
