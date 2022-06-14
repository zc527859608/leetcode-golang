package leetcodeWork

import "fmt"

func TestIsValidSudoku() {
	s := [][]byte{{'.', '.', '.', '.', '5', '.', '.', '1', '.'}, {'.', '4', '.', '3', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '3', '.', '.', '1'}, {'8', '.', '.', '.', '.', '.', '.', '2', '.'}, {'.', '.', '2', '.', '7', '.', '.', '.', '.'}, {'.', '1', '5', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '2', '.', '.', '.'}, {'.', '2', '.', '9', '.', '.', '.', '.', '.'}, {'.', '.', '4', '.', '.', '.', '.', '.', '.'}}
	for _, x := range s {
		for _, y := range x {
			fmt.Print(string(y), "  ")
		}
		fmt.Print("\n")
	}

	println(IsValidSudoku(s))
}

func IsValidSudoku(board [][]byte) bool {
	for x := 0; x < 9; x++ { //每行每列判断是否有重复数字
		countx := [9]int{} //行 数字计数
		county := [9]int{} //列 数字计数
		for y := 0; y < 9; y++ {
			if board[x][y] != '.' {
				if countx[int(board[x][y])-49] == 1 {
					return false
				} else {
					countx[int(board[x][y])-49] = 1
				}
			}
			if board[y][x] != '.' {
				if county[int(board[y][x])-49] == 1 {
					return false
				} else {
					county[int(board[y][x])-49] = 1
				}
			}
		}
	}
	for x := 0; x < 3; x++ { //每个3*3判断
		for y := 0; y < 3; y++ {
			counts := [9]int{}
			for a := x * 3; a < x*3+3; a++ {
				for b := y * 3; b < y*3+3; b++ {
					if board[a][b] != '.' {
						if counts[int(board[a][b])-49] == 1 {
							return false
						} else {
							counts[int(board[a][b])-49] = 1
						}
					}
				}
			}
		}
	}
	return true
}
