package main

import (
	"strings"
)

func solution(str, ending string) bool {
	return len(str) >= len(ending) && str[len(str) - len(ending):] == ending
}

func solution1(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}