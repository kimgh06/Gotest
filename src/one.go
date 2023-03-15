package main

import "fmt"

func main() {
	a := 0
	b := 0
	sum := 0
	fmt.Scan(&a, &b)
	osum := a * b
	for i := 0; i < 3; i++ {
		sum = a * (b % times(10, i+1)) / times(10, i)
		println(sum)
		b = b - b%times(10, i+1)
	}
	println(osum)
}

func times(n int, k int) int {
	sum := 1
	for i := 0; i < k; i++ {
		sum *= n
	}
	return sum
}
