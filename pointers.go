package main

import "fmt"

func main() {
	i := 10
	p := &i
	*p = 20
	fmt.Println("i =", i)
}
