package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var n, m int
var list [7]int
var visit [7]bool

func main() {
	fmt.Fscan(reader, &n, &m)
	recursion(0)
}

func recursion(sidx int) {
	defer writer.Flush()
	if sidx == m {
		for i := 0; i < m; i++ {
			fmt.Fprintf(writer, "%d ", list[i]+1)
		}
		fmt.Fprintf(writer, "\n")
		return
	}
	for i := 0; i < n; i++ {
		if !visit[i] {
			// visit[i] = true
			list[sidx] = i
			// recursion(sidx)
			recursion(sidx + 1)
			// visit[i] = false
		}
	}
}
