package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var slice = primes[1:4]
	fmt.Println(slice)
	slice[1] = 17
	fmt.Println(slice) // 3 17 7

	s := []struct {
		b bool
		i int
	}{
		{true, 1},
		{false, 2},
		{true, 3}}
	fmt.Println(s)

	numbers := make([]int, 5, 5)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
	}
	fmt.Println(numbers)
	numbers = append(numbers, 5, 6, 7)
	fmt.Println(numbers)

	pow := []int {1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy, dy)
	for i, _ := range s {
		s[i] = make([]uint8, dx, dx)
		for j, _ := range s[i] {
			s[i][j] = uint8((i + j) / 2)
		}
	}
	return s
}
