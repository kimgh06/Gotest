package main

import (
	"fmt"
)

var n, m int
var list [8]int
var visit [8]bool

func main() {
	fmt.Scan(&n, &m)
	recursion(0)
}

func recursion(sidx int) {
	if sidx == m {
		var check bool
		for i := 0; i < m-1; i++ {
			if list[i] > list[i+1] {
				check = true
			}
		}
		if check {
			return
		}
		for i := 0; i < m; i++ {
			fmt.Printf("%d ", list[i]+1)
		}
		fmt.Println("")
		return
	}
	for i := 0; i < n; i++ {
		if !visit[i] {
			visit[i] = true
			list[sidx] = i
			recursion(sidx + 1)
			visit[i] = false
		}
	}
}
