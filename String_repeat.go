package main
import (
	"fmt"
)

func main () {
	fmt.Println(RepeatStr(4, "Salom"))
}

func RepeatStr(repetitions int, value string)(s string) {
	for i := 0; i < repetitions; i ++ {
		s += value
	}
	return s
}