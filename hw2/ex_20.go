package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt has a negative value: %f", e)
}

func Sqrt(x float64) (float64, error) {
	z := float64(1)
	s := float64(0)

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	for i := 1; i <= 5; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
	}
	return s, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("Standard answer:", math.Sqrt(2))

	fmt.Println(Sqrt(-2))
	fmt.Println("Standard answer:", math.Sqrt(-2))
}
