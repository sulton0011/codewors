package main

import (
	"fmt"
)

func main () {
	fmt.Println(Points([]string{"1:1","2:2","3:3","4:4","2:2","3:3","4:4","3:3","4:4","4:4"}))
}
func Points(games []string) (sum int) {
	for _, i := range games {
		if i[0] > i[2]{
			sum += 3
		}else if i[0] == i[2] {
			sum += 1
		}
	}
	return sum
}