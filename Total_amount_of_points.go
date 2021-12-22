package main

func Points(games []string) (sum int) {
	for _, i := range games {
		if i[0] > i[2]{
			sum += 3
		}else if i[0] == i[2] {
			sum += 1
		}
	}
	return sum
}