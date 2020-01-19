package top

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func TestLRUCache(t *testing.T) {

	lru := Constructor(10)

	f := bufio.NewScanner(os.Stdin)

	for {
		f.Scan()
		input := strings.Split(f.Text(), ",")
		if len(input) == 2 {
			k, _ := strconv.Atoi(input[0])
			v, _ := strconv.Atoi(input[1])
			lru.Put(k, v)
		} else {
			k, _ := strconv.Atoi(input[0])
			lru.Get(k)
		}

		fmt.Println(Obj2String(lru))
	}

}

func Obj2String(data interface{}) string {

	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}

func TestPermute(t *testing.T) {
	ans := permute([]int{1, 2, 3, 4})

	fmt.Println(ans)
}

func TestDeepCopy(t *testing.T) {
	a := []int{1, 2, 3}
	var b []int
	if err := deepCopy(&b, a); err != nil {
		fmt.Println(err)
		return
	}
	a[0] = 2

	fmt.Println(b)
}

func TestMaxProduct(t *testing.T) {
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}
