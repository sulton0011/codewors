package main

import (
	"fmt"
)

func main () {
	fmt.Println(Perimeter(5))
}

func Perimeter(n int) int {
	sum := 0
	c, b := 1, 1
	for i := 1; i < n; i ++{
		d := b + c;
		c = b;
		b = d;
		sum += d
	}
	return (sum + 2) * 4
}