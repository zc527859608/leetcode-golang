package leetcodeWork

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, inx, middle, right := 0, -1, -1, len(nums)-1
	if nums[left] == target {
		inx = left
	} else if nums[right] == target {
		inx = right
	} else {
		for right > left+1 {
			middle = (left + right) / 2
			if nums[middle] > target {
				right = middle
			} else if nums[middle] < target {
				left = middle
			} else {
				inx = middle
				break
			}
		}
	}
	if inx == -1 {
		return []int{-1, -1}
	}
	return []int{findboundary(nums, inx, -1, target), findboundary(nums, inx, 1, target)}
}

func findboundary(nums []int, index int, offset int, target int) int {
	if (offset == 1 && (index+1 == len(nums) || nums[index+1] > target)) || (offset == -1 && (index == 0 || nums[index-1] < target)) {
		return index //边界上或者下一位都不是target ，即退出条件
	}

	if index+offset < 0 || index+offset >= len(nums) || nums[index+offset] < target || nums[index+offset] > target {
		var off int
		if offset > 0 {
			off = 1
		} else {
			off = -1
		}
		return findboundary(nums, index, off, target) //越界或者不等于target则重置off
	} else { //标准情况下 将off*2  加快遍历速度
		return findboundary(nums, index+offset, offset<<1, target)
	}
	return 0
}
