package main

import (
	"fmt"
)

func main1() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", a+b)
		}
	}
}

func main2() {
	n := 0
	ans := 0

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := 0
			fmt.Scan(&x)
			ans = ans + x
		}
	}
	fmt.Printf("%d\n", ans)
}
