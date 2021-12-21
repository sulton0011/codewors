package main

import (
	"fmt"
)

func main () {
	fmt.Println(OddCount(15023))
}

func OddCount(n int) int{
	return n >> 1
  }