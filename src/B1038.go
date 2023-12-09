package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	N := 0
	fmt.Scanf("%d", &N)
	fmt.Printf("%d\n", cal(N))
}

func cal(N int) int {
	if N/10 == 0 {
		return N
	}
	n := 10
	for i := n + 1; i <= N; { //11~N
		// fmt.Printf("%d\n", n)
		check, digit := check(n)
		if check {
			i++
			n++
		} else {
			n = (n/int(math.Pow(10, float64(digit))) + 1) * int(math.Pow(10, float64(digit)))
		}
	}
	if n > 1000000 {
		return -1
	}
	return n
}

func check(n int) (bool, int) {
	digit := int(math.Floor(math.Log10(float64(n))))
	max := string(strconv.Itoa(n)[0])
	for i := digit - 1; i >= 0; i-- {
		if string(strconv.Itoa(n)[digit-i]) >= max {
			return false, digit - i
		}
	}
	return true, -1
}
