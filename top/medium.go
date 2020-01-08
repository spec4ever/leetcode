package top

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
