package top

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 0}
	m, n := 1, 1
	nums2 := []int{1}

	merge(nums1, m, nums2, n)

	fmt.Println(nums1)

}

func TestEmptySlice(t *testing.T) {
	a := []int{1, 2, 3}

	fmt.Println(a[:0])
}

func TestIsHappy(t *testing.T) {
	fmt.Println(isHappy(19))
}

func TestCountPrimes(t *testing.T) {
	fmt.Println(countPrimes(10))
}
