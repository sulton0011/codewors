package main
import (
	"strings"
)

func High(s string) string {
	eng := "abcdefghijklmnopqrstuvwxyz"
	sLst := strings.Split(s, " ")
	a := ""
	aSum := 0

	for i := range sLst {
		sum := 0
		for j := range sLst[i] {
			sum += strings.Index(eng, string(sLst[i][j])) + 1
		}
		if aSum < sum {
			aSum = sum
			a = sLst[i]
		}
	}
	return a
}