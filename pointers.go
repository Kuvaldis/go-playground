package main

import "fmt"

func main19() {
	i := 10
	p := &i
	*p = 20
	fmt.Println("i =", i)
}
