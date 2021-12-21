package main

import (
	"fmt"
	// "strings"
)

func main () {
	fmt.Println(solution("abc", "c"))
}

func solution(str, ending string) bool {
	return len(str) >= len(ending) && str[len(str) - len(ending):] == ending
}

// func solution(str, ending string) bool {
// 	return strings.HasSuffix(str, ending)
// }

// var solution = strings.HasSuffix
