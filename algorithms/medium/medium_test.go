package medium

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	a := initListNode(9, 8)
	b := initListNode(1)

	t.Log(getListNode(a))
	t.Log(getListNode(b))

	c := addTwoNumbers(a, b)
	t.Log(getListNode(c))
}

func Test_myPow(t *testing.T) {
	type arg struct {
		x      float64
		n      int
		Wanted float64
	}

	testCases := []arg{
		{
			x:      2,
			n:      0,
			Wanted: 1,
		},
		{
			x:      2,
			n:      1,
			Wanted: 2,
		},
		{
			x:      2,
			n:      10,
			Wanted: 1024,
		},
	}

	for _, tc := range testCases {
		result := myPow(tc.x, tc.n)
		if tc.Wanted != result {
			t.Errorf("x is:%f, n is:%d, result:%v  but wanted:%v", tc.x, tc.n, result, tc.Wanted)
		}
	}
}
