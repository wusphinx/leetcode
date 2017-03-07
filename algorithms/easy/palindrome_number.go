package easy

import (
	"fmt"
)

/*
9. Palindrome Number
Determine whether an integer is a palindrome. Do this without extra space.
*/
func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
