// Write a function that takes a string of parentheses, and determines if the order of the parentheses is valid. 
//The function should return true if the string is valid, and false if it's invalid.

// Examples

//  "()"              =>  true
//  ")(()))"          =>  false
//  "("               =>  false
//  "(())((()())())"  =>  true

package main

func ValidParentheses(parens string) bool {
	n := 0

    for _, char := range parens {
      switch char {
      case '(': n++
      case ')': n--
      }
      if n < 0 {return false}
    }
    
    return n == 0
}