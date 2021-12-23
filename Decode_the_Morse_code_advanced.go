package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println(DecodeBits("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"))
}

func DecodeBits(bits string) string {
	bits = strings.ReplaceAll(bits, "00000000000000", "   ")
	bits = strings.ReplaceAll(bits, "000000", " ")
	bits = strings.ReplaceAll(bits, "111111", "-")
	bits = strings.ReplaceAll(bits, "11", ".")
	return DecodeMorse(bits)
}
  
func DecodeMorse(morseCode string) (decoded string) {
	for _, word := range strings.Split(strings.TrimSpace(morseCode), "   ") {
		for _, char := range strings.Split(word, " "){
		  decoded += char
		}
		decoded += " "
	}
	return strings.TrimSpace(decoded)
}