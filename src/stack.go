package main

import (
	"fmt"
)

func main() {
	top := -1
	var list [10]int
	empty := func() bool {
		return top == -1
	}

	full := func() bool {
		return top > 9
	}

	push := func(value int) bool {
		if full() {
			return false
		}
		top++
		list[top] = value
		return true
	}
	pop := func() bool {
		if empty() {
			return false
		}
		list[top] = 0
		top--
		return true
	}
	push(5)
	push(5)
	push(5)
	pop()

	fmt.Printf("%d", list)
}
