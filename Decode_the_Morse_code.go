package main

import (
	"fmt"
	"strings"
)
func main () {
	fmt.Println(DecodeMorse(".... . -.--   .--- ..- -.. ."))
}

func DecodeMorse(morseCode string) (decoded string) {
	fmt.Println(strings.Split(strings.TrimSpace(morseCode), "   "))
	for _, word := range strings.Split(strings.TrimSpace(morseCode), "   ") {
		for _, char := range strings.Split(word, " "){
			fmt.Println(char)
		  decoded += char
		}
		decoded += " "
	  }
	  return strings.TrimSpace(decoded)
	
}
  
  