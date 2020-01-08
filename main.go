package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"leetcode/top"
	"os"
	"strconv"
	"strings"
)

func main() {

	lru := top.Constructor(10)

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
