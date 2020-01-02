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
