package main

import "fmt"

/***
打印字符，输入行数
        1***
    3***    2***
4***    5***    6***
上一行比下一行缩进4个空格，奇数行正序，偶数行倒序，每个数字占4位补*,间隔同样为4个空格
*/

var blacks = "    "

func PrintLine(num int) {
	var blackNum = num - 1 //4个空格的个数
	all := (1 + num) * num / 2
	allnum := make([]int, all+2)
	for i := 0; i < all; i++ {
		allnum[i] = i + 1
	}

	index := 0
	for line := 1; line <= num; line++ {
		for i := 0; i < blackNum; i++ {
			fmt.Print(blacks)
		}
		printNums(index, line, allnum)
		index += line
		blackNum--
		fmt.Print("\n")
	}

}

func printNums(left, line int, nums []int) {
	if line%2 != 0 {
		for i := left; i < left+line; i++ {
			printNum(nums[i])
			if i != left+line-1 {
				fmt.Print(blacks)
			}
		}
	} else {
		for i := left + line - 1; i >= left; i-- {
			printNum(nums[i])
			if i != left {
				fmt.Print(blacks)
			}
		}
	}
}

func printNum(n int) {
	fmt.Print(n)
	if n/10 == 0 {
		fmt.Print("*")

	}
	if n/100 == 0 {
		fmt.Print("*")

	}
	if n/1000 == 0 {
		fmt.Print("*")

	}
}
