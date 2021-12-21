package main

func NbYear(p0 int, percent float64, aug int, p int) int {
	var count int
	a0 := float64(p0)
	a := float64(p)
	au := float64(aug)

	for a0 < a {
		count ++
		a0 += (a0 / 100) * percent
		if a0 >= a {
			break
		}
		a0 += au
	} 

	if aug < 0 {
		return count + 1
	}else {
		return count
	}

}
func NbYear1(p0 int, percent float64, aug int, p int) int {
	if p0 >= p {
		return 0
	  }
	
	  return 1 + NbYear(p0 + int(float64(p0) * percent/100) + aug, percent, aug, p)
	
}