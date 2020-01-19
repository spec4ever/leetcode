package top

import (
	"bytes"
	"encoding/gob"
	"sort"
)

/*
No: 300
思路： dp三板斧：状态定义（包含初始状态），状态转移，返回值
dp[i] = max(dp[j]+1, dp[i]) for j in [0,i)
*/
func lengthOfLIS(nums []int) int {

	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	lengthMax := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		lengthMax[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && lengthMax[j]+1 > lengthMax[i] {
				lengthMax[i] = lengthMax[j] + 1
			}
		}
	}

	result := 1
	for _, v := range lengthMax {
		if v > result {
			result = v
		}
	}

	return result
}

/*
No: 5
*/
func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	res := string(s[0])
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && (j-i+1) > len(res) {
				if s[i:j+1] == reverseString(s[i:j+1]) {
					res = s[i : j+1]
				}
			}
		}
	}

	return res
}

func reverseString(s string) string {

	reverseByte := make([]byte, 0)

	for i := len(s) - 1; i >= 0; i-- {
		reverseByte = append(reverseByte, byte(s[i]))
	}

	return string(reverseByte)
}

/*
No: 146
*/
type CacheNode struct {
	Key int
	Val int
	Pre *CacheNode
	Nxt *CacheNode
}

type LRUCache struct {
	HashMap map[int]*CacheNode
	Head    *CacheNode
	Tail    *CacheNode
	Cap     int
}

func Constructor(capacity int) LRUCache {
	if capacity < 1 {
		return LRUCache{}
	}

	head := &CacheNode{-1, -1, nil, nil}
	tail := &CacheNode{-1, -1, nil, nil}

	head.Nxt = tail
	tail.Pre = head
	hashMap := make(map[int]*CacheNode, capacity)

	return LRUCache{hashMap, head, tail, capacity}
}

func (this *LRUCache) Get(key int) int {
	//key存在: 放尾部
	if val, ok := this.HashMap[key]; ok {
		val.Pre.Nxt = val.Nxt
		val.Nxt.Pre = val.Pre

		//放尾部
		this.Tail.Pre.Nxt = val
		val.Pre = this.Tail.Pre
		val.Nxt = this.Tail
		this.Tail.Pre = val

		return val.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {

	//key存在
	if val, ok := this.HashMap[key]; ok {
		val.Pre.Nxt = val.Nxt
		val.Nxt.Pre = val.Pre
		//放尾部
		this.Tail.Pre.Nxt = val
		val.Pre = this.Tail.Pre
		val.Nxt = this.Tail
		this.Tail.Pre = val
		//更新值
		val.Val = value

		return
	}

	//key不存在：溢出删头部
	if len(this.HashMap) >= this.Cap {
		toBeDelNode := this.Head.Nxt
		this.Head.Nxt = toBeDelNode.Nxt
		toBeDelNode.Nxt.Pre = this.Head
		delete(this.HashMap, toBeDelNode.Key)
	}

	//建节点
	newNode := &CacheNode{key, value, nil, nil}
	//放尾部
	this.Tail.Pre.Nxt = newNode
	newNode.Pre = this.Tail.Pre
	newNode.Nxt = this.Tail
	this.Tail.Pre = newNode
	this.HashMap[key] = newNode

	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/*
No: 46
思路： 回溯三板斧：路径，选择列表，结束条件； 选择->递归回溯->撤销选择
*/

func permute(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}

	traceBack(nums, path, &ans)

	return ans
}

func traceBack(nums []int, path []int, ans *[][]int) {
	//叶子节点，终止
	if len(path) == len(nums) {
		pathAppend := &[]int{}
		deepCopy(pathAppend, path)
		*ans = append(*ans, *pathAppend)
		return
	}

	//选择列表遍历
	for _, v := range nums {
		if !inSlice(v, path) {
			//选择->递归回溯->撤销选择
			path = append(path, v)
			traceBack(nums, path, ans)
			path = path[:len(path)-1]
		}
	}
}

func inSlice(num int, nums []int) bool {
	for _, v := range nums {
		if num == v {
			return true
		}
	}

	return false
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

/*
No: 279 完全平方数
dp[i] = min{dp[i - j * j] for j in [0,sqrt(i))}
*/
func numSquares(n int) int {

	dp := []int{}
	dp = append(dp, []int{0, 1}...)

	for i := 2; i < n+1; i++ {
		dp = append(dp, i)
		for j := 1; j*j <= i; j++ {
			if dp[i-j*j]+1 < dp[i] {
				dp[i] = dp[i-j*j] + 1
			}
		}

	}

	return dp[n]

}

/*
No: 152 乘积最大子序列
*/
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxRes := 1
	tmpRes, tmpResMinus := 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if tmpRes < 0 {
				maxRes = 0
			} else {
				maxRes = tmpRes
			}
			tmpRes = 1
			tmpResMinus = 1
		} else if nums[i] < 0 {
			tmpResMinus = nums[i] * tmpResMinus
			tmpRes = tmpRes * nums[i]
			if tmpResMinus > maxRes {
				maxRes = tmpResMinus
				tmpRes = tmpResMinus
			} else {
				if tmpRes > maxRes {
					maxRes = tmpRes
				}
				tmpRes = 1
			}
		} else {
			tmpRes *= nums[i]
			tmpResMinus *= nums[i]
			if tmpRes > maxRes {
				maxRes = tmpRes
			}
		}
	}

	return maxRes
}

/*
No: 158  排序链表
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {

	valList := []int{}
	cur := head
	for cur != nil {
		valList = append(valList, cur.Val)
		cur = cur.Next
	}

	cur = head
	sort.Ints(valList)
	i := 0

	for cur != nil {
		cur.Val = valList[i]
		cur = cur.Next
		i++
	}

	return head
}
