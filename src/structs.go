package main

import "fmt"

type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{x: 1}
	v3 = Vertex{}
	v4 = &Vertex{1, 2}
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.x = 4
	fmt.Println(v.x)

	p := &v
	p.x = 1e9
	fmt.Println(v)

	fmt.Println(v1, v4, v2, v3)
}
