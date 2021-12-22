package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main () {
	fmt.Println(("Hey fellow warriors"))
}

func SpinWords(str string) string {
	strList := strings.Split(str, " ")
	str = ""
	for i := 0; i < len(strList); i ++{
		if utf8.RuneCountInString(string(strList[i])) >= 5 {
			for j := range strList[i]{
				str += string(strList[i][j])
			}
		}else {
			str += string(strList[i])
		}
	} 
	return str
}// SpinWords