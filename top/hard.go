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

/*
No: 149 直线上最多的点
 [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
*/
func maxPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	if len(points) <= 2 {
		return 1
	}

	ans := 1
	hashMap := map[string]int{} //key为最大公约数斜率string

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			kSharp := twoPointsK(points[j], points[i])
			hashMap[kSharp]++
		}
	}

	for _, v := range hashMap {
		if v > ans {
			ans = v
		}
	}

	return ans
}

func twoPointsK(a, b []int) string {
	deltaY := b[1] - a[1]
	deltaX := b[0] - a[0]

	return string(deltaY + deltaX)
}

func gcdTwoNum(a, b int) int {

	for a != 0 {
		a, b = b%a, a
	}

	return b
}
