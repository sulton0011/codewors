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
	sList := strings.Split(s, " ")
	for i := range sList {
		fmt.Println(i, sList[i], a)
	}
	return ""
}