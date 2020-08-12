package main

import "fmt"

type Vertex struct {
	x, y int
}

var (
	v1 = Vertex{}
	v2 = Vertex{x: 1}
	p1 = &Vertex{1, 2}
)

func main() {
	v := Vertex{1, 2}
	v.x = 4
	fmt.Println(v)

	p := &v
	p.x = 16
	fmt.Println(v)

	fmt.Println(v1, v2, p1)
}
