package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)
	maps := make([]int, n)
	for i := 0; i < n; i++ {
		maps[i] = i + 1
	}
	updown := 0
	for maps[0] != 0 {
		if updown == 0 {
			for i := 0; i < n-1; i++ {
				maps[i] = maps[i+1]
			}
			updown = 1
			continue
		}
		max := 0
		for i = 0; i < n; i++ {
			if maps[i] == 0 {
				break
			}
		}
		updown = 0
	}
}
