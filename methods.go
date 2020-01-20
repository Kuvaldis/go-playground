package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Abs() float64 {
	if v == nil {
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Abser interface {
	Abs() float64
}

type IPAddr [4]byte
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	v := Vertex2{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	var a Abser
	a = &v
	describe(a)
	fmt.Println(a.Abs())
	a = f
	describe(a)
	fmt.Println(a.Abs())

	var v2 *Vertex2
	a = v2
	describe(a)
	fmt.Println(a.Abs())

	var i interface{}
	describe2(i)
	i = 42
	describe2(i)
	i = "hello"
	describe2(i)

	s, ok := i.(string)
	fmt.Println(s, ok)
	fl, ok := i.(float64)
	fmt.Println(fl, ok)

	do(21)
	do("hello")
	do(true)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func describe(a Abser) {
	fmt.Printf("(%v, %T)\n", a, a)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}