package main

import (
	"fmt"
)

func main() {
	// n := 0
	// fmt.Scan(&n)
	maps := make([][]int, 9)
	for i := 0; i < 9; i++ {
		maps[i] = make([]int, 9)
	}
	max := 0
	index_x := 0
	index_y := 0

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Scan(&maps[i][j])
			if maps[i][j] > max {
				max = maps[i][j]
				index_x = i
				index_y = j
			}
		}
	}
	fmt.Printf("%d\n%d %d", max, index_x+1, index_y+1)
}
