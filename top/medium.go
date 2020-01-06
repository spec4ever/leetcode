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
type LRUCache struct {
	CacheKV  map[int]int
	Capacity int
	LRUQueue []int
	LRUMap   map[int]int
}

func Constructor(capacity int) LRUCache {

	newCache := LRUCache{}
	if capacity <= 0 {
		return newCache
	}

	newCache.CacheKV = make(map[int]int, capacity)
	newCache.Capacity = capacity
	newCache.LRUQueue = make([]int, 0)
	newCache.LRUMap = make(map[int]int, 0)

	return newCache
}

func (this *LRUCache) Get(key int) int {

	if val, ok := this.CacheKV[key]; ok {
		//更新队列和index
		index := this.LRUMap[key]
		for i := index + 1; i < len(this.LRUQueue); i++ {
			this.LRUMap[this.LRUQueue[i]]--
		}
		this.LRUQueue = append(this.LRUQueue[0:index], this.LRUQueue[index+1:]...)
		this.LRUQueue = append(this.LRUQueue, key)
		this.LRUMap[key] = len(this.LRUQueue) - 1

		return val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {

	//key存在，更新值、队列、index
	if _, ok := this.CacheKV[key]; ok {
		index := this.LRUMap[key]
		for i := index + 1; i < len(this.LRUQueue); i++ {
			this.LRUMap[this.LRUQueue[i]]--
		}
		this.LRUQueue = append(this.LRUQueue[0:index], this.LRUQueue[index+1:]...)
		this.LRUQueue = append(this.LRUQueue, key)
		this.LRUMap[key] = len(this.LRUQueue) - 1

		this.CacheKV[key] = value
		return
	}

	//key不存在，缓存未满，直接存入
	if len(this.CacheKV) < this.Capacity {
		this.CacheKV[key] = value
		this.LRUQueue = append(this.LRUQueue, key)
	} else {
		//缓存已满： 删除最老key及queue及map
		deleteKey := this.LRUQueue[0]
		delete(this.CacheKV, deleteKey)
		this.CacheKV[key] = value

		for i := 0; i < len(this.LRUQueue); i++ {
			this.LRUMap[this.LRUQueue[i]]--
		}
		this.LRUQueue = this.LRUQueue[1:]
		this.LRUQueue = append(this.LRUQueue, key)

		delete(this.LRUMap, deleteKey)
	}

	//更新index
	this.LRUMap[key] = len(this.LRUQueue) - 1
	return
}
