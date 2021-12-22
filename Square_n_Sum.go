package main
import (
	"fmt"
)
func main () {
	fmt.Println(SquareSum([]int{0, 3, 4, 5}))
}

func SquareSum(numbers []int) (sum int) {
	for i := range numbers {
		sum += numbers[i] * numbers[i]
	}
    return 
}