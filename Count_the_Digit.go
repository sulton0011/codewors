package main

import (
	"fmt"
)

func main () {
	fmt.Println(NbDig(5755, 0))
}

func NbDig(n int, d int) int {
	var count int
	for i := 0; i <= n; i ++ {
		a := i * i
		for a > 0 {
			if a % 10 == d {
				count ++
			}
			a /= 10
		}
	}
	if d == 0 {
		return count + 1
	}else {
		return count
	}
}