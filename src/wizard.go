package main

import (
	"fmt"
)

func main() {
	n := 1
	fmt.Scanf("%d ", &n)
	maps := make([][]int, n)
	for i := 0; i < n; i++ {
		maps[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&maps[i][j])
		}
	}

	times := 0
	fmt.Printf("times")
	fmt.Scan(&times)
	for i := 0; i < times; i++ {
		row := 0
		col := 0
		ran := 0
		fmt.Printf("row col ran:\n")
		fmt.Scan(&row, &col, &ran)
		for j := -1 * ran; j <= ran; j++ {
			if j == 0 {
				if maps[row][col] == 0 {
					maps[row][col] = 1
				} else {
					maps[row][col] *= 2
				}
			} else {
				if row+j >= 0 && row+j < n {
					if maps[row+j][col] == 0 {
						maps[row+j][col] = 1
					} else {
						maps[row+j][col] *= 2
					}
				}
				if col+j >= 0 && col+j < n {
					if maps[row][col+j] == 0 {
						maps[row][col+j] = 1
					} else {
						maps[row][col+j] *= 2
					}
				}
			}
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += maps[i][j]
		}
	}

	fmt.Printf("%d %d", maps, sum)
}
