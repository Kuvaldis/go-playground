package main

import (
	"fmt"
	"strings"
)

type Coordinate struct {
	Lat, Long float64
}

func main16() {
	m1 := map[string]Coordinate{
		"Bell Labs": Coordinate{40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m1)

	m2 := make(map[string]int)
	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])
	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])
	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])
	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		m[word] += 1
	}
	return m
}
