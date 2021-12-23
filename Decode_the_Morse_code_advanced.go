package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println(DecodeBits("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"))
}

func DecodeBits(bits string) string {
	bits = strings.ReplaceAll(bits, "111111", "-")
	bits = strings.ReplaceAll(bits, "11", ".")
	bits = strings.ReplaceAll(bits, "00000000000000", "   ")
	bits = strings.ReplaceAll(bits, "000000", " ")
	bits = strings.ReplaceAll(bits, "0", "")
	return DecodeMorse(bits)
}
  
func DecodeMorse(morseCode string) (decoded string) {
	MORSE_CODE := map[string]string{
		".-" : "A", "-..." : "B", "-.-." : "C",
		"-.." : "D", "." : "E", "..-." : "F",
		"--." : "G", "...." : "H", ".." : "I",
		".---" : "J", "-.-" : "K", ".-.." : "L",
		"--" : "M", "-." : "N", "---" : "O", ".--." : "P",
		"--.-" : "Q", ".-." : "R", "..." : "S", "-" : "T",
		"..-" : "U", "...-" : "V", ".--" : "W", "-..-" : "X",
		"-.--" : "Y", "--.." : "Z", "-----" : "0", ".----" : "1",
		"..---" : "2", "...--" : "3", "....-" : "4", "....." : "5",
		"-...." : "6", "--..." : "7", "---.." : "8", "----." : "9",
	}
	for _, word := range strings.Split(morseCode, "   ") {
		for _, char := range strings.Split(word, " "){
		  decoded += MORSE_CODE[char]
		}
		decoded += " "
	}
	return strings.TrimSpace(decoded)
}