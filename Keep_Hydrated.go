package main

import (
	"fmt"
)

func main () {
	fmt.Println(Litres(11.8))
}

func Litres(time float64) int {
	return int(time) / 2
  }