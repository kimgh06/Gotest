package main

import (
	"fmt"
)

var list [10000]int

func main() {
	fmt.Print(fibo(10))
}

func fibo(n int) int {
	if n < 2 {
		return n
	}
	if list[n] != 0 {
		return list[n]
	}
	list[n] = fibo(n-1) + fibo(n-2)
	return list[n]
}
