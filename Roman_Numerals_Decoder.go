package main
import (
	"fmt"
)
func main () {
	fmt.Println(Decode("XXI"))
}
func Decode(roman string) (sum int) {
	RimSon := map[string]int{
		"I" : 1,
		"V" : 5,
		"X" : 10,
		"L" : 50,
		"C" : 100,
		"D" : 500,
		"M" : 1000,
	}
	for i := len(roman) - 1; i >= 0; i -- {
		sum += RimSon[string(roman[i])]
	}

	return 
}