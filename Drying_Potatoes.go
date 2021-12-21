package main

import (
	"fmt"
)

func main () {
	fmt.Println(Potatoes(82, 127, 80))
}

func Potatoes(p0, w0, p1 int) int {
    return w0 * (100 - p0) / (100 - p1)
}