package main

import (
	"fmt"
	"strconv"

)

func main () {
	fmt.Println(Points([]string{"1:1","2:2","3:3","4:4","2:2","3:3","4:4","3:3","4:4","4:4"}))
}
func Points(games []string) (sum int) {
	for i := range games {
		s, _ := strconv.Atoi(string(games[i][0]))
		a, _ := strconv.Atoi(string(games[i][2]))
		if s > a {
			sum += 3
		}else if s < a {
			sum += 0
		}else {
			sum += 1
		}
	}
	return sum
}