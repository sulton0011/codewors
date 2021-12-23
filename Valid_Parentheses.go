// Write a function that takes a string of parentheses, and determines if the order of the parentheses is valid. 
//The function should return true if the string is valid, and false if it's invalid.

// Examples

//  "()"              =>  true
//  ")(()))"          =>  false
//  "("               =>  false
//  "(())((()())())"  =>  true

package main

import (
	"fmt"
)

func main (){
	fmt.Println(ValidParentheses("()"))
}

func ValidParentheses(parens string) bool {
    count1 := 0
	count := 0
	for _, i := range parens {
		if string(i) == "(" && count1 < count{
			count ++
		}
		if string(i) == ")" && count1 <= count {
			count1 ++
		}
	}
	if count == count1 {
		return true
	}else {
		return false
	}
}