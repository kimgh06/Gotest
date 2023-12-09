package main

import (
	"fmt"
	"time"
)

func main() {
	go yes()
	go yes()
	go yes()
	go yes()
	go yes()
	time.Sleep(time.Millisecond * 2)
}

func yes() {
	fmt.Printf("wow\n")
}
