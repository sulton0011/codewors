package main

func GrowingPlan(upSpeed, downSpeed, desiredHeight int) int {
	var i int = 0
	var count int = 0

	for i < desiredHeight {
		if i >= desiredHeight {
			return count - 1
		}else {
			count ++
			i += upSpeed
			if i >= desiredHeight {
				return count
			}
			i -= downSpeed
		}
	}

	return count - 1
}