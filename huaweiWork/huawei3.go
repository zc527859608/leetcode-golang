package main

import "fmt"

/***
输入两个数，人数与命令行数
后续输入对应行数命令，构成为3个数字，前两个数字标识人，
后一个数字为0标识两人为一组，为1则判断两人是否是一组
依次执行命令，对于命令为1判断为同组输出we are a team
非同组输出we are not a team
非预期输入da pian zi
*/

func main() {
	a := 0
	b := 0
	n, _ := fmt.Scanln(&a, &b)
	if n == 0 {
		fmt.Print("NULL")
		return
	}
	if a < 1 || a > 100000 || b < 1 || b > 10000 {
		fmt.Print("NULL")
		return
	}
	all := make([][]int, b)
	var (
		i1 int
		i2 int
		i3 int
	)
	for i := 0; i < b; i++ {
		num, _ := fmt.Scanln(&i1, &i2, &i3)
		if num == 0 {
			fmt.Print("NULL")
			return
		}
		all[i] = []int{i1, i2, i3}
	}
	runFunc(a, b, all)
}

var out1 = "we are a team"
var out2 = "we are not a team"
var out3 = "da pian zi"

func runFunc(mNum, lNum int, all [][]int) {
	allRes := make([]map[int]int, mNum)
	for i := range allRes {
		allRes[i] = make(map[int]int)
	}
	var (
		m1 int
		m2 int
		l  int
	)
	for i := range all {
		m1 = all[i][0]
		m2 = all[i][1]
		l = all[i][2]
		if m1 < 1 || m1 > mNum || m2 < 1 || m2 > mNum {
			fmt.Printf("%s\n", out3)
			continue
		}
		if l == 0 {
			allRes[m1-1][m2] = 1
			allRes[m2-1][m1] = 1
			continue
		}
		if l == 1 {
			if ok := checkIn(m1, m2, allRes); ok {
				fmt.Printf("%s\n", out1)
			} else {
				fmt.Printf("%s\n", out2)
			}
			continue
		}
		fmt.Printf("%s\n", out3)
	}
}

func checkIn(a, b int, all []map[int]int) bool {
	checked := make(map[int]int)
	checked[a] = 1
	for k := range all[a-1] {
		if checkOne(k, b, all, checked) {
			return true
		}
	}
	return false
}

func checkOne(a, b int, all []map[int]int, checked map[int]int) bool {
	if a == b {
		return true
	}
	checked[a] = 1
	for k := range all[a-1] {
		if _, ok := checked[k]; ok {
			continue
		}
		if checkOne(k, b, all, checked) {
			return true
		}
	}
	delete(checked, a)
	return false
}
