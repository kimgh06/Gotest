package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var n, m int
var list [7]int
var visit [7]bool
var newlist [7]int

var lagacylist [100000][7]int

var location int = 0

var temp [7]int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &newlist[i])
	}
	sort.Ints(newlist[0:n])
	recursion(0)
}

func recursion(sidx int) {
	if sidx == m {
		for i := 0; i < m; i++ {
			temp[i] = newlist[list[i]]
		}
		for i := location; i >= 0; i-- {
			if lagacylist[i] == temp {
				return
			}
		}
		for i := 0; i < m; i++ {
			fmt.Fprintf(writer, "%d ", temp[i])
		}
		lagacylist[location] = temp
		location++
		fmt.Fprintf(writer, "\n")
		return
	}
	for i := 0; i < n; i++ {
		// if !visit[i] {
		// 	visit[i] = true
		list[sidx] = i
		recursion(sidx + 1)
		// 	visit[i] = false
		// }
	}
}
