package main
import (
	"strings"
)

func FindShort(s string) int {
	sLenMin := 30 

	ls := 0
	for _, i := range s {
		if i == 32{
			if sLenMin > ls{
				sLenMin = ls
			}
			ls = 0
		}else {
			ls ++
		}
	}
	if sLenMin > ls{
		sLenMin = ls
	}
	return sLenMin
}

func FindShort1(s string) int {
	c := len(s)
	for _, word := range strings.Fields(s) {
		if c == 0 || len(word) < c {
			c = len(word)
		}
	}
	return c
}

func FindShort2(s string) int {
	shortest := len(s)
	for _, word := range strings.Split(s, " ") {
	  if len(word) < shortest {
		shortest = len(word)
	  }
	}
	return shortest
}