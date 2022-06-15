package main

import "fmt"

/***
对于以元音'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'开始结束的字符串算瑕疵串，其中包含的非元音字符个数为
其瑕疵度，输入一个瑕疵度与一串字符串，返回该字符串中符合瑕疵度要求的最长子串长度
*/

func main() {
	var a = 0
	var b = "asdbuiodevauufgh"
	fmt.Println(findMaxLen(a, b))
}

func findMaxLen(mnum int, str string) int {
	rs := []rune(str)
	dp := make([][]int, len(rs)+1)
	var (
		ostrt []int //每个起始的元音index数组
		max   int
	)
	var (
		left      int  //当前字符的index
		leftIndex int  //当前串瑕疵计算ostrt中起始index
		last      bool //前一个字符是不是元音
	)

	for {
		dp[left] = []int{0, 0, 0}
		if checkO(rs[left]) {
			ostrt = append(ostrt, left)
			last = true
			break
		}
	}

	var (
		nowMaxNum int //最大的瑕疵度（允许范围内）
		nowMinNum int //最小瑕疵度
		nowMaxLen int //最大瑕疵度对应长度
		nowNum    int //前一个连续的非元音字符个数
	)
	for i := left + 1; i < len(rs); i++ {
		if checkO(rs[i]) { //当前是元音
			if last { //前一个字符也是元音
				nowMaxNum = dp[i-1][0]
				nowMinNum = dp[i-1][1]
				nowMaxLen = dp[i-1][2] + 1
			} else {
				nowMaxNum = dp[i-1][0]
				nowMinNum = dp[i-1][1]
				nowMaxLen = dp[i-1][2] + 1
				if nowMaxNum > mnum { //超过瑕疵度限制
					if nowMinNum > mnum { //当前最小瑕疵度串超限，及前序无法构成符合规则的瑕疵串
						nowMaxNum = 0
						nowMaxLen = 1
					} else {
						for index := leftIndex + 1; index < len(ostrt); index++ { //依次后移起点元音字符，找到最大合适瑕疵串
							nowMaxNum -= dp[ostrt[leftIndex]][1]
							if nowMaxNum <= mnum {
								nowMaxLen -= i - ostrt[leftIndex] + 1
								break
							}
						}
					}
				}
			}
			ostrt = append(ostrt, i)
			last = true
			if max < nowMaxLen {
				max = nowMaxLen
			}
		} else { //当前非原音
			if last {
				nowMaxNum = dp[i-1][0] + 1
				nowMinNum = 1
				nowMaxLen = dp[i-1][2] + 1
			} else {
				nowMaxNum = dp[i-1][0] + 1
				nowMinNum = dp[i-1][1] + 1
				nowMaxLen = dp[i-1][2] + 1
			}
			nowNum++
			last = false
		}
		dp[i] = []int{nowMaxNum, nowMinNum, nowMaxLen}
	}

	return max
}

func checkO(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
