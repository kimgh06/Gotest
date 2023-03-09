package main

import (
	"fmt"
)

// 백준 1021번
func main() {
	var max int //최대 크기
	var n int   // 횟수
	fmt.Scan(&max, &n)
	p := 1 // 위치 정보
	var arr []int
	sum := 0
	distance := 0
	rotation := 0
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		if p == num { //동일할 때
			if !contains(arr, p+1) {
				sum++
			}
			p = p%max + 1
		} else if (num-p+max)%max > (p-num+max)%max { // 왼쪽이 더 가까울때
			distance = (p - num + max) % max
			for j := 0; j < distance; j++ {
				if !contains(arr, p-1) {
					sum++
				}
				p = (p + max - 1) % max
			}
			p = p%max + 1
		} else if (num-p+max)%max < (p-num+max)%max { // 오른쪽이 더 가까울때
			distance = (num - p + max) % max
			for j := 0; j < distance; j++ {
				if !contains(arr, p+1) {
					sum++
				}

			}
		} else { // 거리가 동일 할때
			rotation = 2
			distance = (num - p + max) % max
		}

		for j := 0; j < distance; j++ {
			switch rotation {
			case 0:
				p = (num)%max + 1
				if !contains(arr, p+1) {
					sum++
				}
			case 1:
				if !contains(arr, p+1) {

				}
			}
			arr = append(arr, num)

		}
	}
}

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
