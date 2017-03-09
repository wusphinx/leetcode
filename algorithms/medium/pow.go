package medium

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	var pow int
	var result float64
	if n > 0 {
		pow = n
	} else {
		pow = -n
	}

	if n%2 == 0 {
		result = myPow(x*x, pow/2)
	} else {
		result = x * myPow(x*x, pow/2)
	}
	if n < 0 {
		result = 1 / result
	}

	return result
}
