package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println(SpinWords("Hey fellow warriors"))
}

func SpinWords(str string) string {
	s := strings.Split(str, " ")
  
	for i, v := range s {
	  if len(v) >= 5 {
		res := ""
		for _, r := range v {
		  res = string(r) + res
		}
		s[i] = res
	  }
	} 
	
	return strings.Join(s, " ")
}// SpinWords