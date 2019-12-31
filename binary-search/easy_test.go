package binary_search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{13}
	target := 13

	fmt.Println(search(nums, target))
}
