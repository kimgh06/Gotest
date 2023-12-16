package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var a, b big.Int
	fmt.Fscan(reader, &a, &b)
	result := new(big.Int).Add(&a, &b)
	fmt.Fprint(writer, result)
}
