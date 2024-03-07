package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var newlist [500][500]int
var n, m int

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &n, &m)
	max_set := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &newlist[i][j])
			if newlist[i][j] > max_set {
				max_set = newlist[i][j]
			}
		}
	}
	sum_max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if newlist[i][j] == max_set {
				initialSlices := [][]int{{i, j}}
				sum_value := recursion(4, initialSlices, max_set)
				if sum_value > sum_max {
					sum_max = sum_value
				}
			}
		}
	}
	fmt.Println(sum_max)
}

func recursion(cnt int, slices [][]int, sum int) int {
	if cnt == 1 {
		return sum
	}
	find_max := 0
	where_max := []int{0, 0}
	for i := 0; i < len(slices); i++ {
		x := slices[i][0]
		y := slices[i][1]
		if x != 0 && newlist[x-1][y] > find_max && check(x-1, y, slices) {
			find_max = newlist[x-1][y]
			where_max[0] = x - 1
			where_max[1] = y
		}
		if x != n-1 && newlist[x+1][y] > find_max && check(x+1, y, slices) {
			find_max = newlist[x+1][y]
			where_max[0] = x + 1
			where_max[1] = y
		}
		if y != 0 && newlist[x][y-1] > find_max && check(x, y-1, slices) {
			find_max = newlist[x][y-1]
			where_max[0] = x
			where_max[1] = y - 1
		}
		if y != m-1 && newlist[x][y+1] > find_max && check(x, y+1, slices) {
			find_max = newlist[x][y+1]
			where_max[0] = x
			where_max[1] = y + 1
		}
	}
	slices = append(slices, where_max)
	return recursion(cnt-1, slices, sum+find_max)
}

func check(x int, y int, slices [][]int) bool {
	for i := 0; i < len(slices); i++ {
		if slices[i][0] == x && slices[i][1] == y {
			return false
		}
	}
	return true
}
