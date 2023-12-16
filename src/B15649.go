package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int
var list [8]int
var visit [8]bool

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	recursion(0)
}

func recursion(sidx int) {
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	if sidx == m {
		// var check bool
		// for i := 0; i < m-1; i++ {
		// 	if list[i] > list[i+1] {
		// 		check = true
		// 	}
		// }
		// if check {
		// 	return
		// }
		for i := 0; i < m; i++ {
			fmt.Fprintf(writer, "%d ", list[i]+1)
		}
		fmt.Fprintf(writer, "\n")
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
