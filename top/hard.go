package top

/*
No: 42 接雨水
*/
func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	sum, occupy := 0, 0

	index := 0
	//找到第一个不为0的
	for height[index] == 0 {
		index++
	}

	//left表示左边计算起点 [0,1,0,2,1,0,1,3,2,1,2,1]
	left := index
	for index < len(height)-1 {
		index++
		//不低于left，计算积水，更新left，清空occupy
		if height[index] >= height[left] {
			tmpSum := (index - left - 1) * height[left]
			sum = sum + (tmpSum - occupy)
			left = index
			occupy = 0
		} else {
			//低于left，计算occupy
			occupy += height[index]
		}
	}
	//如果遍历完没有不低于left的，对称一下递归计算
	if occupy != 0 {
		reverseHeight := []int{}
		for index >= left {
			reverseHeight = append(reverseHeight, height[index])
			index--
		}
		sum += trap(reverseHeight)
	}

	return sum
}
