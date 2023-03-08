package main

func main() {
	k := 10
	print(fac(k))
}

func fac(k int) int {
	sum := 1
	n := 1
	for n <= k {
		sum *= n
		n++
	}
	return sum
}
