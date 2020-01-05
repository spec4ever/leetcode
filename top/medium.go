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
