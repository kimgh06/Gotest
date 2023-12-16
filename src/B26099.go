package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	fmt.Fscan(reader, &n)

	defer writer.Flush()
	cnt := 0

	for i := 0; i <= n; i += 3 {
		if (n-i)%5 == 0 {
			cnt = i/3 + (n-i)/5
			break
		}
	}

	if cnt == 0 {
		fmt.Print(-1)
		return
	}
	fmt.Fprintf(writer, "%d", cnt)
}
