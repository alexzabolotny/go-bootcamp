package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt is here
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

// Sqrt is here
func Sqrt(s float64) (float64, error) {
	if s < 0 {
		return 0, ErrNegativeSqrt(s)
	}

	return math.Sqrt(s), nil
}

func main() {

}
