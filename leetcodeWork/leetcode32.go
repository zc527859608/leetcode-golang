package leetcodeWork

func LongestValidParentheses(s string) int {
	rs := []rune(s)
	var count, le, max, bre int //len 当前连续有效长度 count 计数标志(0即构成有效括号) bre 下次循环继续节点前一位(跳过相关区域用)
	for i := 0; i < len(rs); i++ {
		if rs[i] == ')' { //由'（'开始匹配
			continue
		}
		count, le, bre = 0, 0, 0 //开始匹配重置
		for x := i; x < len(rs); x++ {
			if rs[x] == '(' {
				count++
			} else {
				count--
			}
			if count == 0 {
				le = x - i + 1 //有效括号长度
				bre = x + 1    //退出循环时应当从最长有效括号的后两位位置继续判断
			} else if count < 0 { //当前为'）'且前面为有效括号
				bre = x //此时可直接跳到'）'的后一位继续判断
				break
			}
		}
		if le > max {
			max = le
		}
		if bre > 0 { //判断是否需要跳过
			i = bre
		}
	}
	return max
}

func LongestValidParentheses1(s string) int {
	inByte := []rune(s)
	if len(inByte) < 2 {
		return 0
	}

	dp := make([]int, len(inByte))
	var max int

	dp[0] = 0
	if inByte[0] == '(' && inByte[1] == ')' {
		dp[1] = 2
		max = 2
	} else {
		dp[1] = 0
	}
	var l int

	for i := 2; i < len(inByte); i++ {
		l = 0 //重置
		if inByte[i] == '(' {
			dp = append(dp, 0)
			continue
		} else {
			if inByte[i-1] == '(' { //与前一位构成有效括号
				l = dp[i-2] + 2 //长度为到前两位的dp+2
			} else if i-dp[i-1]-1 >= 0 && inByte[i-dp[i-1]-1] == '(' { // 即((...))
				if i-dp[i-1]-2 >= 0 { //()((...))
					l = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					l = dp[i-1] + 2
				}
			}
			dp[i] = l
			if l > max {
				max = l //更新max
			}
		}
	}
	return max
}
