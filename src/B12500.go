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
	if m > 3 { //ㅡ
		for i := 0; i < n; i++ {
			for j := 0; j < m-3; j++ {
				if newlist[i][j]+newlist[i][j+1]+newlist[i][j+2]+newlist[i][j+3] > sum_max {
					sum_max = newlist[i][j] + newlist[i][j+1] + newlist[i][j+2] + newlist[i][j+3]
				}
			}
		}
	}
	if n > 3 { //ㅣ
		for i := 0; i < m; i++ {
			for j := 0; j < n-3; j++ {
				if newlist[j][i]+newlist[j+1][i]+newlist[j+2][i]+newlist[j+3][i] > sum_max {
					sum_max = newlist[j][i] + newlist[j+1][i] + newlist[j+2][i] + newlist[j+3][i]
				}
			}
		}
	}
	if n > 1 && m > 1 { //ㅁ
		for i := 0; i < n-1; i++ {
			for j := 0; j < m-1; j++ {
				if newlist[i][j]+newlist[i+1][j]+newlist[i][j+1]+newlist[i+1][j+1] > sum_max {
					sum_max = newlist[i][j] + newlist[i+1][j] + newlist[i][j+1] + newlist[i+1][j+1]
				}
			}
		}
	}
	if n > 2 && m > 1 { //ㅓ ㅏ
		for i := 0; i < n-2; i++ {
			for j := 0; j < m-1; j++ {
				if newlist[i][j]+newlist[i+1][j]+newlist[i+2][j]+newlist[i+1][j+1] > sum_max {
					sum_max = newlist[i][j] + newlist[i+1][j] + newlist[i+2][j] + newlist[i+1][j+1]
				}
				if newlist[i][j+1]+newlist[i+1][j+1]+newlist[i+2][j+1]+newlist[i+1][j] > sum_max {
					sum_max = newlist[i][j+1] + newlist[i+1][j+1] + newlist[i+2][j+1] + newlist[i+1][j]
				}
				//ㄱ l
				if newlist[i][j]+newlist[i][j+1]+newlist[i+1][j+1]+newlist[i+2][j+1] > sum_max {
					sum_max = newlist[i][j] + newlist[i][j+1] + newlist[i+1][j+1] + newlist[i+2][j+1]
				}
				//ㄱ l 반대
				if newlist[i][j]+newlist[i+1][j]+newlist[i+2][j]+newlist[i][j+1] > sum_max {
					sum_max = newlist[i][j] + newlist[i+1][j] + newlist[i+2][j] + newlist[i][j+1]
				}
				//ㄴ l
				if newlist[i][j]+newlist[i+1][j]+newlist[i+2][j]+newlist[i+2][j+1] > sum_max {
					sum_max = newlist[i][j] + newlist[i+1][j] + newlist[i+2][j] + newlist[i+2][j+1]
				}
				//ㄴ l 반대
				if newlist[i][j+1]+newlist[i+1][j+1]+newlist[i+2][j+1]+newlist[i+2][j] > sum_max {
					sum_max = newlist[i][j+1] + newlist[i+1][j+1] + newlist[i+2][j+1] + newlist[i+2][j]
				}
				//ㄹ l
				if newlist[i][j]+newlist[i+1][j]+newlist[i+1][j+1]+newlist[i+2][j+1] > sum_max {
					sum_max = newlist[i][j] + newlist[i+1][j] + newlist[i+1][j+1] + newlist[i+2][j+1]
				}
				//ㄹ l 반대
				if newlist[i][j+1]+newlist[i+1][j+1]+newlist[i+1][j]+newlist[i+2][j] > sum_max {
					sum_max = newlist[i][j+1] + newlist[i+1][j+1] + newlist[i+1][j] + newlist[i+2][j]
				}
			}
		}
	}
	if n > 1 && m > 2 { //ㅡ
		for i := 0; i < m-2; i++ {
			for j := 0; j < n-1; j++ {
				// ㄱ ㅡ
				if newlist[j][i]+newlist[j][i+1]+newlist[j][i+2]+newlist[j+1][i+2] > sum_max {
					sum_max = newlist[j][i] + newlist[j][i+1] + newlist[j][i+2] + newlist[j+1][i+2]
				}
				// ㄴ ㅡ
				if newlist[j][i]+newlist[j+1][i]+newlist[j+1][i+1]+newlist[j+1][i+2] > sum_max {
					sum_max = newlist[j][i] + newlist[j+1][i] + newlist[j+1][i+1] + newlist[j+1][i+2]
				}
				// ㄱ ㅡ 반대
				if newlist[j][i]+newlist[j][i+1]+newlist[j][i+2]+newlist[j+1][i] > sum_max {
					sum_max = newlist[j][i] + newlist[j][i+1] + newlist[j][i+2] + newlist[j+1][i]
				}
				// ㄴ ㅡ 반대
				if newlist[j][i+2]+newlist[j+1][i]+newlist[j+1][i+1]+newlist[j+1][i+2] > sum_max {
					sum_max = newlist[j][i+2] + newlist[j+1][i] + newlist[j+1][i+1] + newlist[j+1][i+2]
				}
				//ㅜ ㅗ
				if newlist[j][i]+newlist[j][i+1]+newlist[j][i+2]+newlist[j+1][i+1] > sum_max {
					sum_max = newlist[j][i] + newlist[j][i+1] + newlist[j][i+2] + newlist[j+1][i+1]
				}
				if newlist[j][i+1]+newlist[j+1][i]+newlist[j+1][i+1]+newlist[j+1][i+2] > sum_max {
					sum_max = newlist[j][i+1] + newlist[j+1][i] + newlist[j+1][i+1] + newlist[j+1][i+2]
				}
				//ㄹ ㅡ
				if newlist[j][i]+newlist[j][i+1]+newlist[j+1][i+1]+newlist[j+1][i+2] > sum_max {
					sum_max = newlist[j][i] + newlist[j][i+1] + newlist[j+1][i+1] + newlist[j+1][i+2]
				}
				//ㄹ ㅡ 반대
				if newlist[j][i+1]+newlist[j][i+2]+newlist[j+1][i]+newlist[j+1][i+1] > sum_max {
					sum_max = newlist[j][i+1] + newlist[j][i+2] + newlist[j+1][i] + newlist[j+1][i+1]
				}
			}
		}
		fmt.Println(sum_max)
	}
}
