package main

import (
	"fmt"
	"log"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:  %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	if i, err := Sqrt(2); err != nil {
		log.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := Sqrt(-2); err != nil {
		log.Println(err)
	} else {
		fmt.Println(i)
	}
}
