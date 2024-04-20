package main

import (
	"fmt"
	"ch2/popcount"
)

func main() {
	fmt.Printf("%d => %d\n", 10, popcount.PopCount(10))
	fmt.Printf("%d => %d\n", 1024, popcount.PopCount(1024))
}
