package leetcodeWork

func search(nums []int, target int) int {
	if (nums[0] < nums[len(nums)-1] && (target < nums[0] || target > nums[len(nums)-1])) || //超出范围直节返回
		(nums[0] > nums[len(nums)-1] && target < nums[0] && target > nums[len(nums)-1]) {
		return -1
	} else if target == nums[0] {
		return 0
	} else if target == nums[len(nums)-1] {
		return len(nums) - 1
	}
	left, middle, right := 0, 0, len(nums)-1
	for right > left+1 {
		middle = (left + right) / 2
		if nums[middle] > target {
			if nums[right] < nums[left] && target < nums[0] && nums[middle] > nums[0] { //特殊情况  456123  target 1 middle上为6
				left = middle
			} else { //其他为标准二分
				right = middle
			}
		} else if nums[middle] < target {
			if nums[left] > nums[right] && target > nums[0] && nums[middle] < nums[0] { //特殊情况 45123 target 5 middle上为1
				right = middle
			} else { //其他为标准二分
				left = middle
			}
		} else {
			return middle
		}
	}
	return -1
}
