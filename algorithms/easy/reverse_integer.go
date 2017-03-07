package easy

import (
	"math"
)

/*
7. Reverse Integer
Reverse digits of an integer.
Example1: x = 123, return 321
Example2: x = -123, return -321
*/
func reverse(x int) (r int) {
	iter := x
	for iter != 0 {
		r = r*10 + iter%10
		iter /= 10
	}

	if r > math.MaxInt32 || r < math.MinInt32 {
		return 0
	}

	return
}
