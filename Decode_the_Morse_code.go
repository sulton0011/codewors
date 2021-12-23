package main

import (
	"fmt"
	"strings"
)
func main () {
	fmt.Println(DecodeMorse(".--"))
}

func DecodeMorse(morseCode string) string {
	Morse := map[string]string {
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
	for i, j := range Morse {
		morseCode = strings.ReplaceAll(morseCode, i, j)
	}
	return morseCode
  }
  
  