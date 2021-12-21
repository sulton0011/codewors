package main

import (
	"fmt"
)


func main () {
	fmt.Println(TwoSum([]int{2,7,11,15}, 9))
}

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers); i ++ {
		for j := i + 1; j < len(numbers); j ++{
			if numbers[i] + numbers[j] == target {
				return [2]int{i, j}
			}
		}
	}
    return [2]int{}
}