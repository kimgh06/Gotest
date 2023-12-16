package main

import "fmt"

func main() {
	fmt.Printf("%d", recursion(12))
}

func recursion(n int) int {
	if n == 1 {
		return 1
	}
	return n * recursion(n-1)
}
