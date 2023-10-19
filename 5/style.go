package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

type Vertex2 struct {
	x int
	y string
}

var A [2]Vertex
var B [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var c [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var D [10]string = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~, ", p.name, p.age)
}

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
	fmt.Println("---------------------------------------------------------------------")
	for i, b := range B {
		fmt.Println(i+1, b*2)
	}

	for _, c := range c {
		fmt.Println(c * 2)
	}
	fmt.Println("---------------------------------------------------------------------")
	myMap := make(map[string]int)
	myMap["key"] = 10
	myMap["key2"] = 20
	fmt.Println(myMap)
	fmt.Println(myMap["key"])
	fmt.Println(myMap["key2"])
	fmt.Println("---------------------------------------------------------------------")
	myMap2 := make(map[string]Vertex2)
	myMap2["key1"] = Vertex2{1, "something"}
	myMap2["key2"] = Vertex2{2, "something else"}
	fmt.Println(myMap2)
	fmt.Println(myMap2["key1"])
	fmt.Println(myMap2["key2"])
	fmt.Println("---------------------------------------------------------------------")
	myMap3 := make(map[string]Vertex2)
	myMap3["key1"] = Vertex2{1, "iti"}
	myMap3["key2"] = Vertex2{2, "ni"}
	myMap3["key3"] = Vertex2{3, "san"}
	myMap3["key4"] = Vertex2{4, "yon"}
	myMap3["key5"] = Vertex2{5, "go"}
	for key := range myMap3 {
		fmt.Println(myMap3[key].x)
		fmt.Println(myMap3[key].y)
		// fmt.Println(myMap3[value.y])
	}
	fmt.Println("---------------------------------------------------------------------")
	for key, value := range myMap3 {
		// fmt.Println(myMap3[key].x)
		// fmt.Println(myMap3[key].y)
		fmt.Println(key)
		fmt.Println(value)
	}
	fmt.Println("---------------------------------------------------------------------")
	person := Person{"james", 20}
	person.greeting()
}
