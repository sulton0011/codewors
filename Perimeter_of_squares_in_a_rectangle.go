package main

import (
	"fmt"
)

func main () {
	fmt.Println(Perimeter(5))
}

func Perimeter(n int) int {
	sum := 0
	for i := 1; i <= n; i ++{
		sum += i
	}
	return sum
}