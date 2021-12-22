package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main () {
	fmt.Println(SpinWords("Hey fellow warriors"))
}

func SpinWords(str string) string {
	strList := strings.Split(str, " ")
	str = ""
	for i := 0; i < len(strList); i ++{
		if utf8.RuneCountInString(string(strList[i])) >= 5 {
			str += " "
			for j := utf8.RuneCountInString(string(strList[i])) - 1; j >= 0; j --{
				str += string(strList[i][j])
			}
		}else {
			str += " " + string(strList[i])
		}
	} 
	return str[1:]
}// SpinWords