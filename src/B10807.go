package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)
	maps := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&maps[i])
	}
	num := 0
	fmt.Scan(&num)
	cnt := 0
	for i := 0; i < n; i++ {
		if maps[i] == num {
			cnt++
		}
	}
	fmt.Printf("%d", cnt)
}
