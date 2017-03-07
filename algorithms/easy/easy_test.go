package easy

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type arg struct {
		Param  int
		Wanted bool
	}

	testCases := []arg{
		{
			Param:  1,
			Wanted: true,
		},
		{
			Param:  12,
			Wanted: false,
		},
		{
			Param:  1212,
			Wanted: false,
		},
		{
			Param:  -2147483648,
			Wanted: false,
		},
		{
			Param:  1221,
			Wanted: true,
		},
	}
	for _, tc := range testCases {
		result := isPalindrome(tc.Param)
		if tc.Wanted != result {
			t.Errorf("param is:%d, result:%v  but wanted:%v", tc.Param, result, tc.Wanted)
		}
	}

}

func Test_twoSum(t *testing.T) {
	type arg struct {
		Nums   []int
		Target int
		Wanted int
	}

	testCases := []arg{
		{
			Nums:   []int{3, 2, 4},
			Target: 6,
			Wanted: 2,
		},
		{
			Nums:   []int{3, 2, 4},
			Target: 100,
			Wanted: 0,
		},
	}
	for _, tc := range testCases {
		result := twoSum(tc.Nums, tc.Target)
		if len(result) != tc.Wanted {
			t.Errorf("param is:%+v, result:%v", tc, result)
		}
	}

}

func Test_reverse(t *testing.T) {
	type arg struct {
		Param  int
		Wanted int
	}

	testCases := []arg{
		{
			Param:  123,
			Wanted: 321,
		},
		{
			Param:  -123,
			Wanted: -321,
		},
		{
			Param:  -1073741824,
			Wanted: 0,
		},
		{
			Param:  2147483648,
			Wanted: 0,
		},
	}

	for _, tc := range testCases {
		result := reverse(tc.Param)
		if tc.Wanted != result {
			t.Errorf("param is:%d, result:%v  but wanted:%v", tc.Param, result, tc.Wanted)
		}
	}
}
