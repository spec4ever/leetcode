package dynamic_programming

/*
No: 121
 [7,1,5,3,6,4]
*/
func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}
	min := prices[0]
	ans := 0
	for _, v := range prices {
		//当前最低点
		if v < min {
			min = v
		}

		//当前最大收益
		if v-min > ans {
			ans = v - min
		}
	}

	return ans
}
