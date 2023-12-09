package main

import (
	"fmt"
)

func main() {
	top := -1
	list := "-----------"
	empty := func() bool {
		return top == -1
	}

	full := func() bool {
		return top > 9
	}
	pop := func() bool {
		if empty() {
			return false
		}
		list = list[:top] + "-" + list[top+1:]
		top--
		return true
	}

	check := func(gwalho string) bool {
		if empty() {
			return false
		}
		if (string(list[top]) == "(" && gwalho == ")") || (string(list[top]) == "[" && gwalho == "]") || (string(list[top]) == "{" && gwalho == "}") {
			pop()
			return true
		}
		return false
	}
	push := func(gwalho string) bool {
		if full() {
			return false
		}
		if gwalho == ")" || gwalho == "}" || gwalho == "]" {
			return check(gwalho)
		} else {
			top++
			list = list[:top] + gwalho + list[top+1:]
			return true
		}
	}

	str := ""
	fmt.Scan(&str)
	for i := 0; i < len(str); i++ {
		if !push(string(str[i])) {
			fmt.Printf("errrr")
			return
		}
	}

	fmt.Printf("no err")
	// fmt.Printf("%d", list)
}
