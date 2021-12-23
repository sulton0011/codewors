package main

func Decode(roman string) int {
	numerals := map[rune]int{
	  'I': 1,
	  'V': 5,
	  'X': 10,
	  'L': 50,
	  'C': 100,
	  'D': 500,
	  'M': 1000,
	}
	
	var result int
	for _, v := range roman {
	  result += numerals[v] - (result % numerals[v])*2
	}
	
	return result
  }