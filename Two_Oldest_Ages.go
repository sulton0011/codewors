package main

func TwoOldestAges(ages []int) [2]int {
	for i:=0; i< len(ages)-1; i++ {
		for j:=0; j < len(ages)-i-1; j++ {
		   if (ages[j] > ages[j+1]) {
			  ages[j], ages[j+1] = ages[j+1], ages[j]
		   }
		}
	 }
	return [2]int{ages[len(ages) - 2], ages[len(ages) - 1]}
}