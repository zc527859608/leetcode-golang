package leetcodeWork

func searchInsert(nums []int, target int) int {
	var left,right int
	right=len(nums)-1

	if nums[left]>=target{  //小于最小元素，插入头部，等于最小元素，返回数组头
		return 0
	} else if nums[right]==target{//等于最大元素。返回数组原尾部
		return right
	}else if nums[right]<target{//大于最小元素，直接加到原数组后面
		return right+1
	}


	for left+1<right{
		middle:=(left+right)/2
		if nums[middle]>target{
			right=middle
		}else if nums[middle]<target {
			left=middle
		}else{
			return middle
		}
	}
	return right
}
