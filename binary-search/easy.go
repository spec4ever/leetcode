package binary_search

//leetcode 704
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
