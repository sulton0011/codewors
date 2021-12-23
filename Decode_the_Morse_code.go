package main

import (
	"strings"
)

func DecodeMorse(morseCode string) (decoded string) {
	for _, word := range strings.Split(strings.TrimSpace(morseCode), "   ") {
		for _, char := range strings.Split(word, " "){
		  decoded += MORSE_CODE[char]
		}
		decoded += " "
	  }
	  return strings.TrimSpace(decoded)	
}
  
  