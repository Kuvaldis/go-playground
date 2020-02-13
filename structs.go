package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func main23() {
	v := Vertex{1, 2}
	v.X = 4
	p := &v
	p.Y = 1e9
	fmt.Println(v)

	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	v3 := Vertex{}
	p = &Vertex{1, 2}
	fmt.Println(v1, v2, v3, p)
}
