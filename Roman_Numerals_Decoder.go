package main

import (
	"fmt"
	// "strings"
)
func main () {
	fmt.Println(Decode("MMVIII"))
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
	for i := 0; i < len(roman); i ++ {
		var a bool
		if i < len(roman) - 2{
			a = (RimSon[string(roman[i + 1])] > RimSon[string(roman[i])] || RimSon[string(roman[i + 2])] > RimSon[string(roman[i])])
			}else if i < len(roman) - 1{
			a = (RimSon[string(roman[i + 1])] > RimSon[string(roman[i])])
		}
		
		if a {
			sum -= RimSon[string(roman[i])]
		}else {
			sum += RimSon[string(roman[i])]
		}
	}

	return sum
}