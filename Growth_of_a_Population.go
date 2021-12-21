package main

import (
	"fmt"
)

func main () {
	fmt.Println(NbYear(1500000, 0.25, -1000, 2000000))
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	// var count int
	// a0 := float64(p0)
	// a := float64(p)
	// au := float64(aug)

	// for a0 < a {
	// 	count ++
	// 	fmt.Println(a0)
	// 	a0 += (a0 / 100) * percent
	// 	if a0 >= a {
	// 		break
	// 	}
	// 	a0 += au

	// } 
	// fmt.Println(a0)
	// if aug < 0 {
	// 	return count + 1
	// }else {
	// 	return count
	// }


	if p0 >= p {
		return 0
	  }
	
	  return 1 + NbYear(p0 + int(float64(p0) * percent/100) + aug, percent, aug, p)
	
}