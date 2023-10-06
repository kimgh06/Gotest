package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := 0
	list := make([]int, 10)
	fmt.Fscan(reader, &n)
	for i := 1; i <= n; i++ {
		maxdigit := len(strconv.Itoa(i))
		power := math.Pow(10, float64(maxdigit))
		for j := int(power); j >= 10; j /= 10 {
			adigit := (i % j) / (j / 10)
			list[adigit]++
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Fprintf(writer, strconv.Itoa(list[i])+" ")
	}
}
