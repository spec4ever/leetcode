package top

import (
	"fmt"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {

	data := []int{1, 2}

	fmt.Println(lengthOfLIS(data))
}

func TestLongestPalindrome(t *testing.T) {
	s := "ac"

	fmt.Println(longestPalindrome(s))
}

func TestReverseString(t *testing.T) {
	s := "baba"

	fmt.Println(reverseString(s))
}
