package main

import (
	"fmt"
)

func main() {
	// n := 0
	// fmt.Scan(&n)
	maps := make([]int, 9)
	max := 0
	index := 0
	for i := 0; i < 9; i++ {
		fmt.Scan(&maps[i])
		if maps[i] > max {
			max = maps[i]
			index = i
		}
	}
	fmt.Printf("%d\n%d", max, index+1)
}
