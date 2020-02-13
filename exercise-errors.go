package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e)) // cannot call just e, sprintf checks interface inside for 'v' verb. if it's error then it calls Error method, you'll get stack overflow
}

func SqrtErr(x float64) (float64, error) {
	if x <= 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for math.Abs(z*z-x) > 0.00001 {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main7() {
	fmt.Println(SqrtErr(2))
	fmt.Println(SqrtErr(-2))
}
