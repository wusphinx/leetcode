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
