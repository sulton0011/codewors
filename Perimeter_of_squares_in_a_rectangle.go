package main

import (
	"fmt"
)

func main () {
	fmt.Println(Perimeter(5))
}

func Perimeter(n int) int {
	if n == 0 {
		return 1
	}
	return n + Perimeter(n - 1)
}