package main
import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println(High("man i need a taxi up to ubud"))
}

func High(s string) string {
	a := ""
	aSum := 0
	sList := strings.Split(s, " ")
	for i := range sList {
		sum := 0
		for j := range sList[i] {
			sum += int(sList[i][j])
		}
		if aSum < sum {
			aSum = sum
			a = sList[i]
		}
	}
	return a
}