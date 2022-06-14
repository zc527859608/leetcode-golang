package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s1 string
	var s2 string
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s1 = input.Text()
	input.Scan()
	s2 = input.Text()
	fmt.Printf("%d\n", getMaxSonLen(s1, s2))
}

func getMaxSonLen(s1, s2 string) int {
	var max int
	rs1 := []rune(s1)
	rs2 := []rune(s2)
	j := len(rs1)
	k := len(rs2)
	dp := make([][]int, j)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	for a := 1; a < j; a++ {
		for b := 1; b < k; b++ {
			if rs1[a] == rs2[b] {
				dp[a][b] = dp[a-1][b-1] + 1
			} else {
				dp[a][b] = 0
			}
			if dp[a][b] > max {
				max = dp[a][b]
			}
		}
	}
	return max
}
