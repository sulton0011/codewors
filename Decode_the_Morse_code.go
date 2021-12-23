package main

import (
	"fmt"
	"strings"
)
func main () {
	// fmt.Println(DecodeMorse(".-- "))
	DecodeMorse(".... . -.--   .--- ..- -.. .")
}

func DecodeMorse(morseCode string) string {
	morseCodeList := strings.Split(morseCode, "   ")
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
	for i, j := range morseCodeList {
		a := strings.Split(j, " ")
		for index, qimat := range a {
			a[index] = strings.ReplaceAll(a[index], qimat, Morse[qimat])
		}
		morseCodeList[i] = strings.Join(a, "")
		fmt.Println(morseCodeList[i])
	}
	return ""
  }
  
  